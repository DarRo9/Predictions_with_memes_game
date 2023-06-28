package main

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
	"math/rand"
)

func main() {
	// Инициализируем GTK.
	gtk.Init(nil)

	// Создаём билдер
	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Загружаем в билдер окно из файла Glade
	err = b.AddFromFile("wdgts.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Получаем объект главного окна по ID
	obj, err := b.GetObject("window_main")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Преобразуем из объекта именно окно типа gtk.Window
	// и соединяем с сигналом "destroy" чтобы можно было закрыть
	// приложение при закрытии окна
	win := obj.(*gtk.Window)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Отображаем все виджеты в окне
	win.ShowAll()

	// Выполняем главный цикл GTK (для отрисовки). Он остановится когда
	// выполнится gtk.MainQuit()

	obj, _ = b.GetObject("entry_1")
	//entry1 := obj.(*gtk.Entry)

	// Получаем кнопку
	obj, _ = b.GetObject("button_1")
	button1 := obj.(*gtk.Button)

	// Настраиваем картинку
	obj, _ = b.GetObject("img")
	img1 := obj.(*gtk.Image)
	img1.SetFromFile("images/11.jpeg")
	list1 := []string{"images/1.jpg", "images/2.jpg", "images/3.jpeg", "images/4.jpg", "images/5.jpg", "images/6.jpg", "images/7.jpg", "images/8.png", "images/9.jpg", "images/10.jpeg"}

	// После нажатия на кнопку меняем картинку с помощью получения псевдорандомного числа
	button1.Connect("clicked", func() {
		num1 := rand.Intn(9)
		img1.SetFromFile(list1[num1])

	})
	gtk.Main()
}
