package main

import (
    "fmt"
    "math"
)

// Функция coffee вычисляет температуру кофе по времени, используя формулу и добавляет результаты в вектор Coffee
func coffee(Tsr, Tc, HeatCoef float64, t int, Coffee *[]float64) {
    for i := 0; i <= t; i++ {
        *Coffee = append(*Coffee, Tc)   // добавляем изначальное значение
        Tc = Tc - HeatCoef*(Tc-Tsr) // высчитываем новое по формуле
    }
}

// Функция aproxA вычисляет коэффициент наклона линии апроксимации
func aproxA(Coffee []float64) float64 {
    var ET, Et, ETt, ET2 float64
    len := len(Coffee)  // количество измерений

    for i := 0; i < len; i++ {
        ET += Coffee[i] // сумма по температуре
        Et += float64(i)    // сумма по времени
        ETt += Coffee[i] * float64(i)   // сумма для произведения по оси температуры и времени
        ET2 += Coffee[i] * Coffee[i]    // сумма для квадрата температуры
    }
    return (float64(len)*ETt - (Et * ET)) / ((float64(len) * ET2) - (ET * ET))
}

// Функция aproxB вычисляет свободный член линии апроксимации
func aproxB(Coffee []float64, a float64) float64 {
    var ET, Et float64
    len := len(Coffee)  // количество измерений

    for i := 0; i < len; i++ {
        ET += Coffee[i] // сумма по температуре
        Et += float64(i)    // сумма по времени
    }
    return (Et - a*ET) / float64(len)
}

// Функция Correl вычисляет коэффициент корреляции между временем и температурой кофе
func Correl(Coffee []float64) float64 {
    var sumTemp float64
    for _, T := range Coffee {
        sumTemp += T
    }

    len := len(Coffee)  // количество измерений
    TMedium := sumTemp / float64(len)   // среднее значение всех измеренных значений
    tMedium := (float64(len) - 1) * float64(len) / 2
    var sumNumbers, tSumSquare, TSumSquare float64

    for i := 0; i < len; i++ {
        sumNumbers += ((float64(i) - tMedium) * (Coffee[i] - TMedium))
        tSumSquare += ((float64(i) - tMedium) * (float64(i) - tMedium))
        TSumSquare += ((Coffee[i] - TMedium) * (Coffee[i] - TMedium))
    }
    return sumNumbers / math.Sqrt(TSumSquare*tSumSquare)
}

// Главная функция программы
func main() {
    fmt.Println("Введите температуру воздуха и кофе, коэффициент охлаждения, максимальное время охлаждения в минутах: ")

    // Ввод данных от пользователя
    var Tsr, Tc, HeatCoef float64   // температура среды, температура кофе, коэффициент охлаждения
    var t int   // время охлаждения
    fmt.Scan(&Tsr, &Tc, &HeatCoef, &t)

    // Вызов функции coffee для расчета температуры кофе по времени
    var Coffee []float64
    coffee(Tsr, Tc, HeatCoef, t, &Coffee)

    // Вывод таблицы температуры кофе по времени
    fmt.Println("Время\tТемпература кофе")
    for i, temperature := range Coffee {
        fmt.Println(i, "\t", fmt.Sprintf("%.2f", temperature))
    }

    // Расчет коэффициента апроксимации и коэффициента корреляции, вывод результатов
    a := aproxA(Coffee) // в a помещаем отклонение по оси времени
    b := aproxB(Coffee, a)  // в b помещаем отклонение по оси температуры
    correl := Correl(Coffee)    // высчитываем коэффициент корреляции
    fmt.Printf("\nЛиния апроксимации: T = %s * t + %s\n", fmt.Sprintf("%.2f", a), fmt.Sprintf("%.2f", b))   // линия апроксимации
    fmt.Printf("\nКоэффициент корреляции %s\n", fmt.Sprintf("%.3f", correl))    //погрешность измерений, нужно для вычисления погрешности
}