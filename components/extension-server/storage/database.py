import os
from typing import List

from sqlmodel.sql.expression import select
from utilities.filepath import get_home_dir
from sqlmodel import create_engine, SQLModel, Session

# these are imported so that the initialization of the database can be done
from schemas.common.event import Event
from schemas.common.extension import Extension


class Singleton(type):
    _instances = {}
    def __call__(cls, *args, **kwargs):
        if cls not in cls._instances:
            cls._instances[cls] = super(Singleton, cls).__call__(*args, **kwargs)
        return cls._instances[cls]


class DB(metaclass=Singleton):
    def __init__(self) -> None:
        self.engine = create_engine("sqlite:///{}".format(os.path.join(get_home_dir(), "extensions.db")))

    def initialize(self) -> None:
        SQLModel.metadata.create_all(self.engine)

    def save(self, obj: SQLModel) -> None:
        session = Session(self.engine)
        session.add(obj)
        session.commit()
        session.refresh(obj)
        return obj

    def fetch_extensions(self)-> List[Extension]:
        with Session(self.engine) as session:
            results = session.exec(select(Extension)).all()
        return results

if __name__=="__main__":
    DB().initialize()