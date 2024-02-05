#include <bits/stdc++.h>
using namespace std;
#define ll int64_t
#include <cstdlib> 
#include <math.h>
#include<cstdio>
#include<cstring>
#include <math.h>
#define vi v(n) vector<ll> v(n);
#define FOR(I, A, B) for (ll I = (A); I <= (B); I++)
#define fo(i,n) for(ll i=0;i<n;i++)
#define sz(a) ll((a).size())
#define pb push_back
#define all(c) (c).begin(),(c).end()
#define dbg(x) cout << #x << " = " << x << endl
#define dbg2(x,y) cout << #x << " = " << x << ", " << #y << " = " << y << endl
#define dbg3(x,y,z) cout << #x << " = " << x << ", " << #y << " = " << y << ", " << #z << " = " << z << endl
#define dbg4(x,y,z,q) cout << #x << " = " << x << ", " << #y << " = " << y << ", " << #z << " = " << z << ", " << #q << " = " << q << endl
#define scan(char_array) scanf("%[^\n]s",&char_array)
#define IOS ios::sync_with_stdio(false); cin.tie(0); cout.tie(0);

// Merges two subarrays of arr[]. 
// First subarray is arr[l..m] 
// Second subarray is arr[m+1..r] 
// Inplace Implementation 

vector<int> arr;

void merge(int start, int mid, int end) 
{ 

	int start2 = mid + 1; 

	// If the direct merge is already sorted 
	if (arr[mid] <= arr[start2]) { 
		return; 
	} 

	// Two pointers to maintain start 
	// of both arrays to merge 
	while (start <= mid && start2 <= end) { 
		

		// If element 1 is in right place 
		if (arr[start] <= arr[start2]) { 
			start++; 
		} 
		else { 
			int value = arr[start2]; 
			int index = start2; 

			// Shift all the elements between element 1 
			// element 2, right by 1. 
			while (index != start) { 
				arr[index] = arr[index - 1]; 
				index--; 
			} 
			arr[start] = value; 

			// Update all the pointers 
			start++; 
			mid++; 
			start2++; 
		} 
	} 
} 

/* l is for left index and r is right index of the 
sub-array of arr to be sorted */
void mergeSort(int l, int r) 
{ 
	if (l < r) { 

		// Same as (l + r) / 2, but avoids overflow 
		// for large l and r 
		int m = l + (r - l) / 2; 

		// Sort first and second halves 
		mergeSort( l, m); 
		mergeSort( m + 1, r); 

		merge( l, m, r); 
	} 
} 

/* UTILITY FUNCTIONS */
/* Function to print an array */

/* Driver program to test above functions */

void populate(int max){
	

for (int i = 0; i < max; i++)
{
	
    arr.push_back(max-i);
}

}
int main() 
{ 
	IOS

	// arr.push_back(12);
	// arr.push_back(11);
	// arr.push_back(13);arr.push_back(5);
	// arr.push_back(6);
	// arr.push_back(71);
	// arr.push_back(17);
	// arr.push_back(37);
	// arr.push_back(20);arr.push_back(72);
	populate(600000);

	 clock_t start, end; 
  
    /* Recording the starting clock tick.*/
    start = clock(); 
  
   //mergeSort( 0, arr.size() - 1); 
  sort(all(arr));
    // Recording the end clock tick. 
    end = clock(); 
  
    // Calculating total time taken by the program. 
    double time_taken = double(end - start) / double(CLOCKS_PER_SEC); 
    cout << "Time taken by program is : " << fixed  
         << time_taken << setprecision(5); 
    cout << " sec " << endl; 
	


	
	return 0; 
} 
