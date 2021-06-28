package main

import (
	"fmt"
	"pomodoro/userSession"
)

// консольный прил помодоро.
// запуск бинарника, запрос имени, запрос задачи, запрос временного интервала.
// счетчик сохраняет количество подходов
// команды старт и стоп

func main() {

	//переменные для сохранения параметров из консоли
	var (
		Name             string
		TimeToWork       int
		SmallTimeToRelax int
		BigTimeToRelax   int
	)

	fmt.Print("Ваше имя: ")
	fmt.Scan(&Name)
	fmt.Print("Укажи время работы в минутах: ")
	fmt.Scan(&TimeToWork)
	fmt.Print("Укажи короткий перерыв: ")
	fmt.Scan(&SmallTimeToRelax)
	fmt.Print("Укажи длинный перерыв: ")
	fmt.Scan(&BigTimeToRelax)

	session := userSession.NewSessions()

	session.SetName(Name).
		SetTimeToWork(TimeToWork).
		SetSmallTimeToRelax(SmallTimeToRelax).
		SetBigTimeToRelax(BigTimeToRelax)

	session.Start(session.GetTimeToWork())

}
