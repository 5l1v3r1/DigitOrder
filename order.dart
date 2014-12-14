typedef cbFunc(List<int> res);

int numberInBase(List<int> number, int base) {
  int res = 0;
  int place = 1;
  for (int i = 0; i < number.length; ++i) {
    res += place * number[number.length - (i + 1)];
    place *= base;
  }
  return res;
}

bool checkCombination(List<int> current, int base) {
  for (int i = 1; i <= current.length; ++i) {
    var l = new List.from(current.sublist(0, i));
    if ((numberInBase(l, base) % i) != 0) {
      return false;
    }
  }
  return true;
}

void bestNumber(List<int> current, List<int> remaining, int base, cbFunc cb) {
  if (!checkCombination(current, base)) {
    return;
  }
  if (remaining.length == 0) {
    cb(current);
    return;
  }
  for (int i = 0; i < remaining.length; i++) {
    var newCur = new List.from(current);
    var newRem = new List.from(remaining);
    newRem.remove(remaining[i]);
    newCur.add(remaining[i]);
    bestNumber(newCur, newRem, base, cb);
  }
}

void main() {
  for (int i = 2; i < 30; i++) {
    print('trying base $i');
    List<int> list = new List();
    for (int j = 0; j < i; ++j) {
      list.add(j);
    }
    bestNumber([], list, i, (list) {
      print('got solution $list');
    });
  }
}
