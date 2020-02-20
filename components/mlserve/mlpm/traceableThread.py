# Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

#coding:utf-8


class ThreadPrinter:
    def __init__(self, filename):
        self.fhs = {}
        self.filename = filename

    def write(self, value):
        try:
            f = self.fhs.get(threading.current_thread().uuid)
            if f is None:
                f = open(self.filename, 'a')
            f.write(value)
            self.fhs[threading.current_thread().uuid] = f
            f.flush()
        except IOError as identifier:
            logging.warn(identifier)


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
        full_log_path = os.path.join(getLogDir(), self.id + ".full.log")
        self.__run_backup = self.run
        self.run = self.__run
        sys.stdout = ThreadPrinter(full_log_path)
        threading.Thread.start(self)

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
