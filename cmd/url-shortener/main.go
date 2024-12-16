package url_shortener

import "fmt"

type Window struct {
	place string `my meta info - it is tag`
}

type Car struct {
	Window
	wheelNum                           int
	passengerOneName, passengerTwoName string `tag attached to both fields`
}

func main() {
	//TODO: init config: cleanenv

	//TODO: init logger: slog

	//TODO: init storage: sqlite

	//TODO: init router: chi "chi render"

	//TODO: run server:

	var myCar = Car{
		Window:           Window{place: "front"},
		wheelNum:         4,
		passengerOneName: "Dima",
		passengerTwoName: "Andrew",
	}
	fmt.Println(myCar)
}
