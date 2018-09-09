import json


def is_jsonable(x):
    try:
        json.dumps(x)
        return True
    except Exception as ex:
        print(ex)
    except:
        return False


class Bundle(object):
    __SOLVERS__ = []

    def add_solver(self, solver):
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
