
import org.springframework.util.StopWatch;

import java.sql.Array;
import java.util.List;

public class MergeSort {

  static Integer[] arr;





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
    static void populate(int max) {
        arr=new Integer[max];


        for (int i = 0; i < max; i++) {

            arr[i]=(max - i);
        }
    }

    public static void main(String[] args) {

        int max=100000;
        MergeSort  mergeSort=new MergeSort();

        populate(max);


        StopWatch watch = new StopWatch();
        watch.start();
        // call to the methods you want to benchmark
        mergeSort.mergeSort( 0, max - 1);
        watch.stop();
        long result = watch.getTotalTimeMillis();
        System.out.println(result);
// some time passes






    }




}
