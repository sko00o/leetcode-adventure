// Forward declaration of guess API.
// @param num, your guess
// @return -1 if my number is lower, 1 if my number is higher, otherwise return
// 0
int guess(int num);

class Solution {
 public:
  int guessNumber(int n) {
    if (n <= 1) {
      return n;
    }

    int low = 1, high = n, mid;
    while (low < high) {
      mid = low + ((high - low) >> 1);
      switch (guess(mid)) {
        case 0:
          return mid;
        case 1:
          low = mid + 1;
          break;
        case -1:
          high = mid - 1;
          break;
      }
    }
    return low;
  }
};