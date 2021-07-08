class Solution:
    def getImportance(self, employees, id):
        """
        :type employees: Employee
        :type id: int
        :rtype: int
        """
        employee_dict = {
            e.id: e for e in employees
        }
        queue = [id]
        traversed = set()
        importance = 0
        while queue:
            id_ = queue.pop(0)
            if id_ in traversed:
                continue
            traversed.add(id_)
            employee = employee_dict[id_]
            importance += employee.importance
            queue.extend(employee.subordinates)
        return importance
