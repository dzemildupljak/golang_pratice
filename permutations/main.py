from copy import copy
from typing import List

class Solution:
    def permute(self, nums):
        res = []
        self.dfs(nums, [], res)
        return res

    def dfs(self, number, path, res):
        if not number:
            res.append(path)
        for i in range(len(number)):
            self.dfs(number[:i]+number[i+1:], path + [number[i]],res)

# class Solution:
#     def permute(self, nums: List[int]) -> List[List[int]]:
#         ln = len(nums)
#         res = list()

#         for i, v in enumerate(nums):
#             # for j in range(self.fakt(ln)):
#             for j in range(ln):
#                 tmp = nums.copy()
#                 tmp.pop(i)
#                 tmp.insert(j, v)
#                 print(tmp)
#                 if tmp in res:
#                     continue
#                 res.append(tmp)

#         print(len(res))
#         print(res)
#         return res

#     def fakt(self, num):
#         res = 1
#         for i in range(1, num+1):
#             res *= i
#         return res

print(Solution().permute([1,2,3, 4]))