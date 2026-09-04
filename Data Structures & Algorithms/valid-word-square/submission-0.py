class Solution:
    def validWordSquare(self, words: List[str]) -> bool:
        for row in range(len(words)):
            for col in range(len(words[row])):
                # The transposed row must exist.
                if col >= len(words):
                    return False

                # The transposed character must exist.
                if row >= len(words[col]):
                    return False

                # Compare words[row][col] with words[col][row].
                if words[row][col] != words[col][row]:
                    return False

        return True