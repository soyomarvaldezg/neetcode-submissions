class Solution:
    def replaceElements(self, arr: List[int]) -> List[int]:
        # Get the length of the input array
        n = len(arr)
        
        # Create a new array of the same length, filled with 0s initially.
        # "ans = [0] * n" creates a list with n copies of 0.
        # Example: [0] * 4 gives [0, 0, 0, 0]
        # We'll overwrite these 0s with the correct values as we iterate.
        ans = [0] * n
        
        # Track the maximum value we've seen to the right.
        # Start at -1 because the last element has nothing to its right,
        # and -1 is what we want to place there (as specified by the problem).
        current_max = -1
        
        # Iterate from the LAST index down to the FIRST (right to left).
        # range(n - 1, -1, -1) means: start at n-1, go down to 0 (inclusive).
        # The -1 step makes it go backwards.
        for i in range(n - 1, -1, -1):
            # Place the current maximum of elements to the right
            # into this position of the answer array.
            # On the first iteration (i = n-1), this is -1.
            ans[i] = current_max
            
            # Update current_max by including arr[i] itself.
            # This way, when we move to position i-1 next,
            # current_max represents the max of everything from i to the end.
            current_max = max(current_max, arr[i])
        
        # Return the completed answer array
        return ans
        