import unittest
import urllib.request
import urllib.error
import urllib.parse
import json
import os
from heartbeat import BaseTest


class Test(BaseTest):
    def __init__(self, *args):
        self.proc = None
        super(Test, self).__init__(*args)

    def test_telemetry(self):
        """
        Test that telemetry metrics are correctly registered and increment / decrement
        """
        # This test is flaky https://github.com/k0ffee/beats/issues/8966
        raise unittest.SkipTest

        if os.name == "nt":
            # This test is currently skipped on windows because file permission
            # configuration isn't implemented on Windows yet
            raise unittest.SkipTest

        server = self.start_server("hello world", 200)
        try:
            self.setup_dynamic(["-E", "http.enabled=true"])

            cfg_file = "test.yml"

            self.write_dyn_config(
                cfg_file, self.http_cfg(
                    "http://localhost:{}".format(server.server_port))
            )

            self.wait_until(lambda: self.output_has(lines=1))

            self.assert_stats({
                "http": {
                    "monitor_starts": 1,
                    "monitor_stops": 0,
                    "endpoint_starts": 1,
                    "endpoint_stops": 0,
                }
            })
            self.assert_state({
                "http": {
                    "monitors": 1,
                    "endpoints": 1,
                }
            })

            tcp_hosts = ["localhost:123", "localhost:456"]

            self.write_dyn_config(
                cfg_file, self.tcp_cfg(*tcp_hosts)
            )

            for tcp_host in tcp_hosts:
                self.wait_until(lambda: self.log_contains(
                    "Start job 'tcp-tcp@{}".format(tcp_host)))

            init_lines = self.output_lines()
            self.wait_until(lambda: self.output_has(lines=init_lines + 2))

            self.assert_stats({
                "http": {
                    "monitor_starts": 1,
                    "monitor_stops": 1,
                    "endpoint_starts": 1,
                    "endpoint_stops": 1,
                },
                "tcp": {
                    "monitor_starts": 1,
                    "monitor_stops": 0,
                    "endpoint_starts": 2,
                    "endpoint_stops": 0,
                }
            })
            self.assert_state({
                "tcp": {
                    "monitors": 1,
                    "endpoints": 2,
                }
            })
        finally:
            self.proc.check_kill_and_wait()
            server.shutdown()

    def assert_state(self, expected={}):
        stats = json.loads(urllib.request.urlopen(
            "http://localhost:5066/state").read())

        total_monitors = 0
        total_endpoints = 0

        for proto in ("http", "tcp", "icmp"):
            proto_expected = expected.get(proto, {})
            monitors = proto_expected.get("monitors", 0)
            endpoints = proto_expected.get("endpoints", 0)
            total_monitors += monitors
            total_endpoints += endpoints
            self.assertDictEqual(stats['heartbeat'][proto], {
                'monitors': monitors,
                'endpoints': endpoints,
            })

        self.assertEqual(stats['heartbeat']['monitors'], total_monitors)
        self.assertEqual(
            stats['heartbeat']['endpoints'], total_endpoints)

    def assert_stats(self, expected={}):
        stats = json.loads(urllib.request.urlopen(
            "http://localhost:5066/stats").read())

        for proto in ("http", "tcp", "icmp"):
            proto_expected = expected.get(proto, {})
            self.assertDictEqual(stats['heartbeat'][proto], {
                'monitor_starts': proto_expected.get("monitor_starts", 0),
                'monitor_stops': proto_expected.get("monitor_stops", 0),
                'endpoint_starts': proto_expected.get("endpoint_starts", 0),
                'endpoint_stops': proto_expected.get("endpoint_stops", 0),
            })
