func lengthOfLongestSubstring(s string) int {

    lastSeen := make(map[byte]int)

    left := 0
    maxLen := 0

    for r :=0;r<len(s);r++{
        ch := s[r]

        if i,exists := lastSeen[ch];exists && i >=left {
            left = i +1
        }

        lastSeen[ch] = r

        currentLen := r - left +1

        if currentLen > maxLen {
            maxLen = currentLen
        } 
    }
    
    return maxLen
}