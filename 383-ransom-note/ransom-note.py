class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        ransomNoteMap = {}
        magazineMap = {}

        for ch in ransomNote:
            ransomNoteMap[ch] = ransomNoteMap.get(ch, 0) + 1
        for ch in magazine:
            magazineMap[ch] = magazineMap.get(ch, 0) + 1
        for key, value in ransomNoteMap.items():
            if key not in magazineMap:
                return False
            if magazineMap[key] < value:
                return False
        return True
        