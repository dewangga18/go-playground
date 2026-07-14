package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// #
func TestRaceCondition(t *testing.T) {
	x := 0

	for range 1000 {
		go func() {
			for range 100 {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Result X:", x)
}

// #
func TestHandleWithMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for range 1000 {
		go func() {
			for range 100 {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Result X:", x)
}

// #
type BankAccount struct {
	mu      sync.RWMutex
	balance int
}

func (account *BankAccount) Deposit(amount int) {
	account.mu.Lock()
	account.balance += amount
	account.mu.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.mu.RLock()
	defer account.mu.RUnlock()
	return account.balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for range 10 {
		go func() {
			for range 1000 {
				account.Deposit(1)
			}
			fmt.Println("Current Balance:", account.GetBalance())
		}()
	}

	time.Sleep(10 * time.Millisecond)
	fmt.Println("Result X:", account.GetBalance())
}

// #
type UserBalance struct {
	sync.Mutex
	name    string
	balance int
}

func (u *UserBalance) Lock() {
	u.Mutex.Lock()
}

func (u *UserBalance) Unlock() {
	u.Mutex.Unlock()
}

func (u *UserBalance) Change(amount int) {
	u.balance = u.balance + amount
}

func TransferDeadlock(to, from *UserBalance, amount int) {
	to.Lock()
	fmt.Println("Lock Increasing", to.name)
	to.Change(amount)

	time.Sleep(2 * time.Second)

	from.Lock()
	fmt.Println("Lock Decreasing", from.name)
	from.Change(-amount)

	time.Sleep(2 * time.Second)

	to.Unlock()
	from.Unlock()

	// output never arrive cause deadlock
	fmt.Println("Unlock", to.name)
	fmt.Println("Unlock", from.name)
}

func TestDealockSimulation(t *testing.T) {
	user1 := UserBalance{name: "Aaron", balance: 500}
	user2 := UserBalance{name: "Evan", balance: 400}

	fmt.Println("Current balance")
	fmt.Println("Aaron :", user1.balance)
	fmt.Println("Evan :", user2.balance)

	go TransferDeadlock(&user2, &user1, 50)
	go TransferDeadlock(&user1, &user2, 35)

	time.Sleep(5 * time.Second)

	fmt.Println()
	fmt.Println("Final balance")
	fmt.Println("Aaron :", user1.balance)
	fmt.Println("Evan :", user2.balance)
}

func TestDealockSimulationWithWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	user1 := UserBalance{name: "Aaron", balance: 500}
	user2 := UserBalance{name: "Evan", balance: 400}

	wg.Add(2)

	fmt.Println("Current balance")
	fmt.Println("Aaron :", user1.balance)
	fmt.Println("Evan :", user2.balance)

	go func() {
		defer wg.Done()
		TransferDeadlock(&user2, &user1, 50)
	}()

	go func() {
		defer wg.Done()
		TransferDeadlock(&user1, &user2, 35)
	}()

	wg.Wait() // wait untill time out and printed panic

	fmt.Println()
	fmt.Println("Final balance")
	fmt.Println("Aaron :", user1.balance)
	fmt.Println("Evan :", user2.balance)
}

// #
func TransferWG(to, from *UserBalance, amount int) {
	fmt.Println("Lock Increasing", to.name)

	if to.name < from.name {
		to.Lock()
		from.Lock()
	} else {
		from.Lock()
		to.Lock()
	}

	defer to.Unlock()
	defer from.Unlock()

	to.Change(amount)
	from.Change(-amount)

	fmt.Println("Unlock", to.name)
}

func TestTransferWithoutDeadlock(t *testing.T) {
	var wg sync.WaitGroup

	user1 := UserBalance{name: "Aaron", balance: 500}
	user2 := UserBalance{name: "Evan", balance: 400}

	fmt.Println("Current balance")
	fmt.Println("Aaron :", user1.balance)
	fmt.Println("Evan :", user2.balance)

	wg.Add(2)

	go func() {
		defer wg.Done()
		TransferWG(&user2, &user1, 50)
	}()

	go func() {
		defer wg.Done()
		TransferWG(&user1, &user2, 35)
	}()

	wg.Wait()

	fmt.Println()
	fmt.Println("Final balance")
	fmt.Println("Aaron :", user1.balance)
	fmt.Println("Evan :", user2.balance)
}

// #
func TestOnce(t *testing.T) {
	var once sync.Once
	var group sync.WaitGroup

	for range 100 {
		group.Add(1)
		go func() {
			defer group.Done()
			once.Do(func() {
				fmt.Println("Loading config...")
				time.Sleep(2 * time.Second)
			})
		}()
	}

	group.Wait()
}

// #
func TestPool(t *testing.T) {
	var pool sync.Pool

	pool.Put("Aaron")
	pool.Put("Evan")
	pool.Put("Juli")

	for range 10 {
		go func() {
			data := pool.Get()
			fmt.Println(data) // some prints will be nil, cause no one is putting anything into the pool after use
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}

func TestPoolNew(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return "New"
		},
	}

	pool.Put("Aaron")
	pool.Put("Evan")
	pool.Put("Juli")

	for range 10 {
		go func() {
			data := pool.Get()
			fmt.Println(data) // no nil prints cause New is defined, so get() will return "New" if there are no data in the pool
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}

type JSONParser struct {
	Data []byte
}

var parserPool = sync.Pool{
	New: func() any {
		fmt.Println("Create parser")
		return &JSONParser{}
	},
}

func Parse(data []byte) {
	parser := parserPool.Get().(*JSONParser)

	parser.Data = data
	fmt.Println(string(parser.Data))
	parser.Data = nil
	parserPool.Put(parser)

}

func TestAnotherPool(t *testing.T) {
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			Parse(fmt.Appendf(nil, "User %d", i))
		}(i)
	}

	wg.Wait()
}

// #

