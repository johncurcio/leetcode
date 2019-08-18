class Solution {
    
    // O(m+n) solution using merge from mergesort
    private int[] merge(int[] nums1, int[] nums2){
        int[] merged = new int[nums1.length + nums2.length];
        
        int i = 0, j = 0, k = 0;
        
        while (i < nums1.length && j < nums2.length){
            if (nums1[i] < nums2[j]){
                merged[k++] = nums1[i++];
            } else {
                merged[k++] = nums2[j++];
            }
        }
        
        while (i < nums1.length){
            merged[k++] = nums1[i++];
        }
        
        while (j < nums2.length){
            merged[k++] = nums2[j++];
        }
        
        return merged;
    }
    
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int[] merged = merge(nums1, nums2);
        int n = merged.length;
        if (n == 0) return 0.0;
        if (n % 2 == 0){
            return (merged[n/2] + merged[n/2-1])/2.0;
        } else {
            return merged[n/2];
        }
    }
}