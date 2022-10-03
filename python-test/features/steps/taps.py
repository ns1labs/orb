from hamcrest import *
import json
from utils import UtilsManager


class Taps(UtilsManager):
    def __init__(self):
        self.taps = dict()

    def __build_tap(self, name, input_type, configs_list, filters_list):
        tap = {
            name: {
                "input_type": input_type,
                "config": {
                },

                "filter": {
                }
            }
        }

        tap = UtilsManager.update_object_with_filters_and_configs(self, tap, name, configs_list, filters_list)

        self.taps.update(tap)

    def add_pcap(self, name, **kwargs):
        self.name = name
        self.pcap_file = {'pcap_file': kwargs.get('pcap_file')}
        self.pcap_source = {'pcap_source': kwargs.get('pcap_source')}
        self.iface = {'iface': kwargs.get('iface')}
        self.host_spec = {'host_spec': kwargs.get('host_spec')}
        self.debug = {'debug': kwargs.get('debug')}
        self.bpf = {'bpf': kwargs.get('bpf')}

        pcap_configs = [self.pcap_file, self.pcap_source, self.iface, self.host_spec, self.debug]

        pcap_filters = [self.bpf]

        self.__build_tap(self.name, "pcap", pcap_configs, pcap_filters)

        return self.taps

    def add_flow(self, name, **kwargs):
        self.name = name
        self.pcap_file = {"pcap_file": kwargs.get('pcap_file')}
        self.port = {"port": kwargs.get('port')}
        self.bind = {"bind": kwargs.get('bind')}
        self.flow_type = {"flow_type": kwargs.get('flow_type')}

        flow_configs = [self.pcap_file, self.port, self.bind, self.flow_type]

        flow_filters = []

        self.__build_tap(self.name, "flow", flow_configs, flow_filters)

        return self.taps

    def add_dnstap(self, name, **kwargs):
        self.name = name
        self.dnstap_file = {"dnstap_file": kwargs.get('dnstap_file')}
        self.socket = {"socket": kwargs.get('socket')}
        self.tcp = {"tcp": kwargs.get('tcp')}
        self.only_hosts = {"only_hosts": kwargs.get('only_hosts')}

        dnstap_configs = [self.dnstap_file, self.socket, self.tcp]

        dnstap_filters = [self.only_hosts]

        self.__build_tap(self.name, "dnstap", dnstap_configs, dnstap_filters)

        return self.taps

    def add_configs(self, name, **kwargs):
        if "config" not in self.taps[name].keys():
            self.taps[name].update({"config": {}})

        self.taps[name]["config"] = UtilsManager.add_configs(self, self.taps[name]["config"], **kwargs)

        return self.taps

    def add_filters(self, name, **kwargs):
        if "filter" not in self.taps[name].keys():
            self.taps[name].update({"filter": {}})

        self.taps[name]["filter"] = UtilsManager.add_filters(self, self.taps[name]["filter"], **kwargs)

        return self.taps

    def add_tag(self, tap_name, tags):
        assert_that(tap_name, any_of(is_in(list(self.taps.keys())), equal_to('all')), "Invalid tap")
        if tap_name == "all":
            tap_names = list(self.taps.keys())
        else:
            tap_names = [tap_name]
        for tap_name in tap_names:
            for tag_key, tag_value in tags.items():
                if "tags" in self.taps[tap_name].keys():
                    self.taps[tap_name]["tags"].update({tag_key: tag_value})
                else:
                    self.taps[tap_name].update({"tags": {tag_key: tag_value}})

        return self.taps

    def remove_tap(self, name):
        assert_that(name, is_in(list(self.taps.keys())), "Invalid tap")
        self.taps.pop(name)

        return self.taps

    def remove_configs(self, name, *args):
        assert_that(name, is_in(list(self.taps.keys())), "Invalid tap")
        self.taps[name]["config"] = UtilsManager.remove_configs(self, self.taps[name]["config"], *args)

        return self.taps

    def remove_filter(self, name, *args):
        assert_that(name, is_in(list(self.taps.keys())), "Invalid tap")
        self.taps[name]["filter"] = UtilsManager.remove_filters(self, self.taps[name]["filter"], *args)

        return self.taps

    def remove_tag(self, tap_name, tags_keys):
        assert_that(tap_name, any_of(is_in(list(self.taps.keys())), equal_to('all')), "Invalid tap")
        if tap_name == "all":
            tap_names = list(self.taps.keys())
        else:
            tap_names = [tap_name]
        for tap_name in tap_names:
            for tag_key in tags_keys:
                if "tags" in self.taps[tap_name].keys():
                    self.taps[tap_name].pop(tag_key, None)

        return self.taps

    def json(self):
        return json.dumps(self.taps)


# class Taps:
#     def __init__(self):
#         pass
#
#     @classmethod
#     def pcap(cls, name, input_type="pcap", settings=None):
#         available_options = ["pcap_source", "iface", "host_spec", "debug", "bpf"]
#         filters_list = ["bpf"]
#
#         return make_tap(name, input_type, available_options, settings, filters_list)
#
#     @classmethod
#     def flow(cls, name, input_type="flow", settings=None):
#         available_options = ["port", "bind", "flow_type"]
#
#         return make_tap(name, input_type, available_options, settings)
#
#     @classmethod
#     def dnstap(cls, name, input_type="dnstap", settings=None):
#         available_options = ["socket", "tcp", "only_hosts"]
#
#         filters_list = ["only_hosts"]
#
#         return make_tap(name, input_type, available_options, settings, filters_list)
#
#
# def make_tap(name, input_type, available_options, settings, filters_list=None):
#     if filters_list is None:
#         filters_list = []
#     kwargs_configs = list(settings.keys())
#
#     assert_that(set(kwargs_configs).issubset(available_options), is_(True),
#                 f"Invalid configuration to tap {input_type}. \n "
#                 f"Options are: {available_options}. \n"
#                 f"Passed: {kwargs_configs}")
#
#     filters = None
#     for tap_filter in filters_list:
#         if tap_filter in kwargs_configs:
#             filters = {tap_filter: settings[tap_filter]}
#             kwargs_configs.remove(tap_filter)
#
#     if len(kwargs_configs) > 0:
#         configs = dict()
#     else:
#         configs = None
#     for configuration in kwargs_configs:
#         configs.update({configuration: settings[configuration]})
#
#     tap = {name: {"input_type": input_type}}
#
#     if filters is not None:
#         filters = {"filter": filters}
#         tap[name].update(filters)
#
#     if configs is not None:
#         configs = {"config": configs}
#         tap[name].update(configs)
#     return tap
