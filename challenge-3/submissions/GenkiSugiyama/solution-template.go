package main

import "fmt"
import "slices"

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type Manager struct {
	Employees []Employee
}

// AddEmployee adds a new employee to the manager's list.
func (m *Manager) AddEmployee(e Employee) {
	// TODO: Implement this method
	m.Employees = append(m.Employees, e)
}

// RemoveEmployee removes an employee by ID from the manager's list.
func (m *Manager) RemoveEmployee(id int) {
	// TODO: Implement this method
	// slicesパッケージのビルトインメソッド
	// DeleteFunc()内のコールバック関数内の条件がtrueを返す要素を削除する
	// →コールバック関数内で要素を削除するための条件を定義する
	m.Employees = slices.DeleteFunc(m.Employees, func(e Employee) bool {
	    return e.ID == id
	})
}

// GetAverageSalary calculates the average salary of all employees.
func (m *Manager) GetAverageSalary() float64 {
	// TODO: Implement this method
	var result float64 
	for _, e := range m.Employees {
	    result += e.Salary
	}
	return result / 2
}

// FindEmployeeByID finds and returns an employee by their ID.
// 「*型名」でその型の値を保持しているポインタ
// Manager構造体内のEmployeesの要素はEmployee型なので&eでEmployeeが格納されている変数のポインタを返してあげる
func (m *Manager) FindEmployeeByID(id int) *Employee {
	// TODO: Implement this method
	for _, e := range m.Employees {
	    if e.ID == id {
	        return &e
	    }
	}
	return nil
}

func main() {
	manager := Manager{}
	manager.AddEmployee(Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
	manager.AddEmployee(Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})
	manager.RemoveEmployee(1)
	averageSalary := manager.GetAverageSalary()
	employee := manager.FindEmployeeByID(2)

	fmt.Printf("Average Salary: %f\n", averageSalary)
	if employee != nil {
		fmt.Printf("Employee found: %+v\n", *employee)
	}
}
