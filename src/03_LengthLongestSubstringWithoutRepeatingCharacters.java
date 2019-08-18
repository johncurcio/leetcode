class Solution {
    public int lengthOfLongestSubstring(String s) {
        int start = 0, end = 0;
        int maxLen = Integer.MIN_VALUE;
        Set<Character> set = new HashSet<Character>();
        while(start < s.length() && end < s.length()){
            if (!set.contains(s.charAt(end))){
                set.add(s.charAt(end++));
            } else {
                set.remove(s.charAt(start++));
            }
            maxLen = Math.max(maxLen, set.size());
        }
        return maxLen == Integer.MIN_VALUE ? 0 : maxLen;
    }
}