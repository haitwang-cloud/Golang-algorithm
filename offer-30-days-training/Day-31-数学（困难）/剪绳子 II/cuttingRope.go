package main

import "fmt"

//链接：https://leetcode.cn/problems/jian-sheng-zi-ii-lcof/solution/mian-shi-ti-14-ii-jian-sheng-zi-iitan-xin-er-fen-f/

func cuttingRope(n int) int {

	if n <= 3 {
		return n - 1
	}
	a, b, p, x, rem := n/3-1, n%3, 1000000007, 3, 1
	for a > 0 {
		if a%2 == 1 {
			rem = (rem * x) % p
		}
		x = x * x % p
		a /= 2
	}
	if b == 0 {
		//3^(a+1) % p
		return (rem * 3) % p
	}
	if b == 1 {
		// 3^a * 4 % p
		return (rem * 4) % p
	}
	return (rem * 6) % p
}

/*


        if n <= 3: return n - 1
        a, b, p, x, rem = n // 3 - 1, n % 3, 1000000007, 3 , 1
        while a > 0:
            if a % 2: rem = (rem * x) % p
            x = x ** 2 % p
            a //= 2
        if b == 0: return (rem * 3) % p # = 3^(a+1) % p
        if b == 1: return (rem * 4) % p # = 3^a * 4 % p
        return (rem * 6) % p # = 3^(a+1) * 2  % p

作者：jyd

*/
func main() {
	test := 10
	fmt.Println(cuttingRope(test))

}
