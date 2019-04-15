import json
from models.config import Config
from models.host import Host


class ConfigManager:
    """ Reads the config.json and populates required data. """
    def __init__(self):
        self.config = self.__read_config()

    def get_hosts(self):
        return self.config.hosts

    def __read_config(self):
        with open('config.json') as config_file:
            config = json.load(config_file)

            # read the entries for Docker hosts
            hosts = []
            for entry in config["hosts"]:
                hosts.append(Host(entry["name"], entry["IP"]))
        return Config(hosts)
