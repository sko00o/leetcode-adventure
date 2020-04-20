// Forward declaration of isBadVersion API.
bool isBadVersion(int version);

class Solution {
 public:
  int firstBadVersion(int n) {
    if (n < 1) {
      return -1;
    }

    int low = 1, high = n;
    while (low < high) {
      int mid = low + ((high - low) >> 1);
      if (isBadVersion(mid)) {
        high = mid;
      } else {
        low = mid + 1;
      }
    }

    if (low <= n && isBadVersion(low)) {
      return low;
    }
    return -1;
  }
};
