# coding:utf-8
import os
import sys
import uuid
import threading
from cvpm.config import getLogDir


class TracedThread(threading.Thread):
    def __init__(self, *args, **keywords):
        threading.Thread.__init__(self, *args, **keywords)
        self.id = None
        self.killed = False
        self.oldout = None

    @property
    def uuid(self):
        return self.id

    def start(self):
        if self.id is None:
            self.id = str(uuid.uuid4)
        full_log_path= os.path.join(getLogDir(), self.id+".log")
        self.oldout = sys.stdout
        with open(full_log_path, 'w') as f:
            sys.stdout = f
            self.__run_backup = self.run
            self.run = self.__run
            threading.Thread.start(self)
            sys.stdout = self.oldout

    def __run(self):
        sys.settrace(self.globaltrace)
        self.__run_backup()
        self.run = self.__run_backup

    def globaltrace(self, frame, event, arg):
        if event == 'call':
            return self.localtrace
        else:
            return None

    def localtrace(self, frame, event, arg):
        if self.killed:
            if event == 'line':
                raise SystemExit()
        return self.localtrace

    def kill(self):
        print('Terminating ' + self.id + '...')
        sys.stdout = self.oldout
        self.killed = True
