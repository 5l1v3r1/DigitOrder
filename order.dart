int numberInBase(List<int> number, int base) {
  int res = 0;
  int place = 1;
  for (int i = 0; i < number.length; ++i) {
    res += place * number[number.length - (i + 1)]
    place *= base;
  }
  return res;
}
