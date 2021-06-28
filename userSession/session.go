package userSession

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type session struct {
	name             string
	task             string
	timeToWork       int
	smallTimeToRelax int
	bigTimeToRelax   int
	counter          int //считает
	isRelax          bool
}

func NewSessions() *session {
	return &session{}
}

func (s *session) SetName(name string) *session {
	s.name = name
	return s
}

func (s *session) GetName() string {
	return s.name
}

func (s *session) SetTask(task string) *session {
	s.task = task
	return s
}

func (s *session) SetCounter(counter int) *session {
	s.counter = counter
	return s
}

func (s *session) SetTimeToWork(timeToWork int) *session {
	s.timeToWork = timeToWork
	return s
}

func (s *session) GetTimeToWork() int {
	return s.timeToWork
}

func (s *session) SetSmallTimeToRelax(smallTimeToRelax int) *session {
	s.smallTimeToRelax = smallTimeToRelax
	return s
}

func (s *session) GetSmallTimeToRelax() int {
	return s.timeToWork
}

func (s *session) SetBigTimeToRelax(bigTimeToRelax int) *session {
	s.bigTimeToRelax = bigTimeToRelax
	return s
}

func (s *session) GetBigTimeToRelax() int {
	return s.timeToWork
}

func (s *session) SetIsRelax(isRelax bool) {
	s.isRelax = isRelax
}

func (s *session) GetisRelax() bool {
	return s.isRelax
}

func (s *session) CreatResponse() {
	strconv.Itoa(s.timeToWork)
	var response string
	switch s.GetisRelax() {
	case false:
		response = "\nСосредоточтесь, " + s.name + ".\nНачало через: "
	case true:
		response = "\nВремя почилить, " + s.name + ".\nНачало через: "
	}
	fmt.Println(response)
	for i := 3; i >= 1; i-- {
		time.Sleep(time.Second)
		if i == 1 {
			fmt.Println(strconv.Itoa(i) + "\nПогнали!\n")
			break
		}
		fmt.Println(i)
	}
}

func (s *session) ChangeParameters() {

	var changeParameters int
	mes := "\nКакие параметры меняем?\n" +
		"1. Изменить имя\n" +
		"2. Изменить время работы\n" +
		"3. Изменить короткий перерыв\n" +
		"4. Изменить длинный перерыв \n" +
		"5. Изменить все параметры\n" +
		"6. Назад\n" +
		"Укажи номер пункта: "
	fmt.Print(mes)
	fmt.Scan(&changeParameters)

	var (
		Name             string
		TimeToWork       int
		SmallTimeToRelax int
		BigTimeToRelax   int
	)

	switch changeParameters {
	case 1:
		fmt.Print("Ваше имя: ")
		fmt.Scan(&Name)
		s.SetName(Name)
		fmt.Print("\nИзменения добавлены!\n")
		s.Choice()

	case 2:
		fmt.Print("Укажи время работы в минутах: ")
		fmt.Scan(&TimeToWork)
		s.SetTimeToWork(TimeToWork)
		fmt.Print("\nИзменения добавлены!\n")
		s.Choice()

	case 3:
		fmt.Print("Укажи время короткого перерыва: ")
		fmt.Scan(&SmallTimeToRelax)
		s.SetTimeToWork(SmallTimeToRelax)
		fmt.Print("\nИзменения добавлены!\n")
		s.Choice()

	case 4:
		fmt.Print("Укажи время длиного перерыва: ")
		fmt.Scan(&BigTimeToRelax)
		s.SetTimeToWork(BigTimeToRelax)
		fmt.Print("\nИзменения добавлены!\n")
		s.Choice()

	case 5:
		fmt.Print("Ваше имя: ")
		fmt.Scan(&Name)
		fmt.Print("Укажи время работы в минутах: ")
		fmt.Scan(&TimeToWork)
		fmt.Print("Укажи время короткого перерыва: ")
		fmt.Scan(&SmallTimeToRelax)
		fmt.Print("Укажи время длинный перерыва: ")
		fmt.Scan(&BigTimeToRelax)
		s.SetName(Name).
			SetTimeToWork(TimeToWork).
			SetSmallTimeToRelax(SmallTimeToRelax).
			SetBigTimeToRelax(BigTimeToRelax)
		fmt.Print("\nИзменения добавлены!\n")
		s.Choice()

	case 6:
		s.Choice()

	default:
		fmt.Print("\nНеверный номер пункта\n")
		s.ChangeParameters()
	}
}

func (s *session) Choice() {
	var choice int
	mes :=
		"\nЧто делаем дальше?\n" +
			"1. Пропустить перерыв\n" +
			"2. Начать короткий перерыв\n" +
			"3. Начать длинный перерыв\n" +
			"4. Изменить параметры\n" +
			"5. Выйти из программы\n" +
			"Укажи номер пункта: "
	fmt.Print(mes)
	fmt.Scan(&choice)

	switch choice {
	case 1:
		s.Start(s.GetTimeToWork())
	case 2:
		s.SetIsRelax(true)
		s.Start(s.GetSmallTimeToRelax())
	case 3:
		s.SetIsRelax(true)
		s.Start(s.GetBigTimeToRelax())
	case 4:
		s.ChangeParameters()
	case 5:
		s.Stop()
	default:
		fmt.Println("Неверный номер пункта")
		s.Choice()
	}
}

// Стартует таймер
func (s *session) Start(timeAny int) {
	// количество символов строки прогресса
	const progressLineCharacters = 60
	s.CreatResponse()
	timer := time.NewTimer(time.Minute * time.Duration(timeAny))
	//	Рисует строку прогресса
	for i := 1; i <= progressLineCharacters; i++ {
		time.Sleep(time.Duration(timeAny) * time.Minute / progressLineCharacters)
		//color.Cyan("#")
		fmt.Print("#")
		if i == progressLineCharacters {
			fmt.Print("\n")
		}
	}

	time.Sleep(time.Microsecond)
	<-timer.C

	switch s.GetisRelax() {
	case true:
		fmt.Print("\nПора за работу!\n")
	case false:
		fmt.Print("\nВремя сделать перерыв!\n")
	}
	s.Choice()
}

func (s *session) Stop() {
	fmt.Printf("\nПрощай, %v \n", s.GetName())
	os.Exit(0)
}
