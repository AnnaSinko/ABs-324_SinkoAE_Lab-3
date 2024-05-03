// Lab 3.3_ABs-324 Sinko.cpp : Этот файл содержит функцию "main". Здесь начинается и заканчивается выполнение программы.
//

#include <iostream>
#include <math.h>
#include <vector>
#include <iomanip>
#include <locale.h>

using namespace std;

// Объявление используемых функций
int coffee(double Tsr, double Tc, double HeatCoef, int t, vector<double>& Coffee);
double aproxA(vector<double> Coffee);
double aproxB(vector<double> Coffee, double a);
double Correl(vector<double> Coffee);

// Главная функция программы
int main() {
    setlocale(LC_ALL, "rus");   //установка русской локали
    system("chcp 65001");
    cout << "Введите температуру воздуха и кофе, коэффициент охлаждения, максимальное время охлаждение в минутах: " << endl;

    // Ввод данных от пользователя: температура среды, температура кофе, коэффициент охлаждения и время охлаждения
    double Tsr, Tc, HeatCoef;   // температура среды, температура кофе, коэффициент охлаждения
    int t;  // время охлаждения
    cin >> Tsr >> Tc >> HeatCoef >> t;

    // Вызов функции coffee для расчета температуры кофе по времени
    vector<double> Coffee;
    coffee(Tsr, Tc, HeatCoef, t, Coffee);

    // Вывод таблицы температуры кофе по времени
    cout << "Время" << "\t" << "Температура кофе" << endl;
    int time = 0;
    for (auto temperature : Coffee) {
        cout << time << "\t" << temperature << setprecision(4) << endl; // вывели все значение от 0 до времени
        time++;
    }

    // Расчет коэффициента апроксимации и коэффициента корреляции, вывод результатов
    double a = aproxA(Coffee);  // в a помещаем отклонение по оси времени
    double b = aproxB(Coffee, a);   // в b помещаем отклонение по оси температуры
    double correl = Correl(Coffee); // высчитываем коэффициент корреляции
    cout << endl << "Линия апроксимации: " << "T = " << a << setprecision(3) << " * t + " << b << setprecision(4) << endl;  // линия апроксимации
    cout << endl << "Коэффициент корреляции " << correl << setprecision(2) << endl; //погрешность измерений, нужно для вычисления погрешности
}

// Функция coffee вычисляет температуру кофе по времени, используя формулу и добавляет результаты в вектор Coffee
int coffee(double Tsr, double Tc, double HeatCoef, int t, vector<double>& Coffee) {
    for (int i = 0; i <= t; i++) {
        Coffee.push_back(Tc);   // добавляем изначальное значение
        Tc = Tc - HeatCoef * (Tc - Tsr);    // высчитываем новое по формуле
    }
    return 0;
}

// Функция aproxA вычисляет коэффициент наклона линии апроксимации
double aproxA(vector<double> Coffee) {

    double ET = 0, Et = 0, ETt = 0, ET2 = 0;
    int len = Coffee.size();    // количество измерений

    for (int i = 0; i < len; i++) {
        ET += Coffee[i];    // сумма по температуре
        Et += i;    // сумма по времени

        ETt += Coffee[i] * i;   // сумма для произведения по оси температуры и времени
        ET2 += Coffee[i] * Coffee[i];   // сумма для квадрата температуры
    }
    return (len * ETt - (Et * ET)) / (len * ET2 - ET * ET);
}

// Функция aproxB вычисляет свободный член линии апроксимации
double aproxB(vector<double> Coffee, double a) {

    double ET = 0, Et = 0;
    int len = Coffee.size();    // количество измерений

    for (int i = 0; i < len; i++) {
        ET += Coffee[i];    // сумма по температуре
        Et += i;    // сумма по времени
    }
    return (Et - a * ET) / len;
}

// Функция Correl вычисляет коэффициент корреляции между временем и температурой кофе
double Correl(vector<double> Coffee) {

    double sumTemp = 0; // сумма температур
    for (double T : Coffee) {
        sumTemp += T;
    }

    int len = Coffee.size();    // количество измерений
    double TMedium = sumTemp / len; // среднее значение всех измеренных значений
    double tMedium = (len - 1) * len / 2;
    double sumNumbers = 0;  // арифметическая сумма значений произведения температуры и времени
    double tSumSquare = 0;  // арифметическая сумма квадрата времени
    double TSumSquare = 0;  // арифметическая сумма квадрата температуры

    for (int i = 0; i < len; i++) {
        sumNumbers += ((i - tMedium) * (Coffee[i] - TMedium));
        tSumSquare += ((i - tMedium) * (i - tMedium));
        TSumSquare += ((Coffee[i] - TMedium) * (Coffee[i] - TMedium));
    }
    return sumNumbers / sqrt(TSumSquare * tSumSquare);
}