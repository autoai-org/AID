import json


def is_jsonable(x):
    """
    Tells you if an object is jsonable
    :return: if it is jsonable or not
    :rtype: bool
    :param x: the object you are testing
    """
    try:
        json.dumps(x)
        return True
    except Exception as ex:
        print(ex)
    except:
        return False


class Bundle(object):
    __SOLVERS__ = []

    def add_solver(self, solver) -> None:
        """
        Adds a solver
        :param self:
        :param solver: solver to add
        :return: None
        """
        self.__SOLVERS__.append(solver)

    @classmethod
    def members(self):
        results = {}
        for attr in dir(self):
            if not attr.startswith("__"):
                if is_jsonable(getattr(self, attr)):
                    results[attr] = getattr(self, attr)
        print(results)
        return results
