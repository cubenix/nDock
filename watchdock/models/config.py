from typing import List
from models.host import Host


class Config:
    """ Represents the object in config.json file. """
    def __init__(self, hosts: List[Host]):
        self.hosts = hosts
        self.host_dict = self.__get_host_dict()

    def __get_host_dict(self):
        host_dict = {}
        for host in self.hosts:
            host_dict[host.name] = host.IP
        return host_dict
