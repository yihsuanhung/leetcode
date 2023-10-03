class Solution:
    def equalPairs(self, grid):
        element_map = {}
        result = 0
        for row in grid:
            row_element = ",".join(map(str, row))
            element_map[row_element] = element_map.get(row_element, 0) + 1
        # prepare col_element
        col_elements = []
        for c in range(len(grid[0])):  # loop through column first
            col_values = []
            for r in range(len(grid)):
                col_values.append(grid[r][c])
            col_element = ",".join(map(str, col_values))
            col_elements.append(col_element)
        # loop through col_elements
        for e in col_elements:
            if str(e) in element_map:
                result += element_map[str(e)]
        return result
