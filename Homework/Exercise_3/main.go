package main

import "fmt"

// function declinationOfNumber(count, one, two, five) {
//     const n100 = Math.abs(count) % 100;
//     const n10 = n100 % 10;

//     if (n100 >= 5 && n100 <= 20) {
//       return five;
//     }
//     if (n10 > 1 && n10 < 5) {
//       return two;
//     }
//     if (n10 === 1) {
//       return one;
//     }
//     return five;
//   }

  
//   const applesNumber = +prompt(`Введите число`);
//   const word = declinationOfNumber(applesNumber, "яблоко", "яблока", "яблок");
//   console.log(`У меня есть ${applesNumber} ${word}.`);

  /*
  1. Получаем остаток от деления на 100, т.к. свыше 100 все повторяется.
  2. Дополнительно будем проверять остаток от деления на 10.
  3. Правила по которым возвращаем слово:
    1. Если последние два числа в диапазоне [5 .. 20], значит слово для числа
      пять.
    2. Если последние два числа вне диапазона [5 .. 20], то все просто:
      2.1. Если единица числа - это число в диапазоне [2 .. 4], значит слово для
        числа два.
      2.2. Если единица числа это число 1, значит слово для числа один.
      2.3. Если единица числа это любое другое число, значит возвращаем слово
        для числа пять.
   */
//






func main() {
	fmt.Println("Start")
}