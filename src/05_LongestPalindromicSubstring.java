class Solution {    
    public String longestPalindrome(String s) {
        if (s == null || s.length() == 0) return "";
        
        int start = 0, end = 0;
        for (int center = 0; center < s.length(); center++){
            int l1 = expandAroundCenter(s, center, center);
            int l2 = expandAroundCenter(s, center, center + 1);
            int l = Math.max(l1, l2);
            if (l > end - start){
                start = center - (l - 1)/2;
                end = start + l - 1;
            }
        }
        
        return s.substring(start, end + 1);
    }
    
    private int expandAroundCenter(String s, int i, int j){
        
        while (i >= 0 && j < s.length() && s.charAt(i) == s.charAt(j)){
            i--;
            j++;
        }

        return j - i - 1;
    }
}