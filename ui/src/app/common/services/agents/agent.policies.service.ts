import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import 'rxjs/add/observable/empty';

import { environment } from 'environments/environment';
import { NotificationsService } from 'app/common/services/notifications/notifications.service';
import { NgxDatabalePageInfo, OrbPagination } from 'app/common/interfaces/orb/pagination.interface';
import { AgentPolicy } from 'app/common/interfaces/orb/agent.policy.interface';

// default filters
const defLimit: number = 20;
const defOrder: string = 'name';
const defDir = 'desc';

@Injectable()
export class AgentPoliciesService {
  paginationCache: any = {};

  cache: OrbPagination<AgentPolicy>;

  backendsCache: OrbPagination<{ [propName: string]: any }>;

  constructor(
    private http: HttpClient,
    private notificationsService: NotificationsService,
  ) {
    this.clean();
  }

  public static getDefaultPagination(): OrbPagination<AgentPolicy> {
    return {
      limit: defLimit,
      order: defOrder,
      dir: defDir,
      offset: 0,
      total: 0,
      data: null,
    };
  }

  clean() {
    this.cache = {
      limit: defLimit,
      offset: 0,
      order: defOrder,
      total: 0,
      dir: defDir,
      data: [],
    };
    this.paginationCache = {};
  }

  addAgentPolicy(agentPolicyItem: AgentPolicy) {
    return this.http.post(environment.agentPoliciesUrl,
        { ...agentPolicyItem },
        { observe: 'response' })
      .map(
        resp => {
          return resp;
        },
      )
      .catch(
        err => {
          this.notificationsService.error('Failed to create Agent Policy',
            `Error: ${ err.status } - ${ err.statusText } - ${ err.error.error }`);
          return Observable.throwError(err);
        },
      );
  }

  getAgentPolicyById(id: string): any {
    return this.http.get(`${ environment.agentPoliciesUrl }/${ id }`)
      .map(
        resp => {
          return resp;
        },
      )
      .catch(
        err => {
          this.notificationsService.error('Failed to fetch Agent Policy',
            `Error: ${ err.status } - ${ err.statusText }`);
          return Observable.throwError(err);
        },
      );
  }

  editAgentPolicy(agentPolicy: AgentPolicy): any {
    return this.http.put(`${ environment.agentPoliciesUrl }/${ agentPolicy.id }`, agentPolicy)
      .map(
        resp => {
          return resp;
        },
      )
      .catch(
        err => {
          this.notificationsService.error('Failed to edit Agent Policy',
            `Error: ${ err.status } - ${ err.statusText }`);
          return Observable.throwError(err);
        },
      );
  }

  deleteAgentPolicy(agentPoliciesId: string) {
    return this.http.delete(`${ environment.agentPoliciesUrl }/${ agentPoliciesId }`)
      .map(
        resp => {
          this.cache.data.splice(this.cache.data.map(ap => ap.id).indexOf(agentPoliciesId), 1);
          return resp;
        },
      )
      .catch(
        err => {
          this.notificationsService.error('Failed to Delete Agent Policies',
            `Error: ${ err.status } - ${ err.statusText }`);
          return Observable.throwError(err);
        },
      );
  }

  getAgentsPolicies(pageInfo: NgxDatabalePageInfo, isFilter = false) {
    const offset = !!pageInfo ? pageInfo.offset : this.cache.offset;
    const limit = pageInfo.limit || this.cache.limit;
    let params = new HttpParams()
      .set('offset', (offset * limit).toString())
      .set('limit', limit.toString())
      .set('order', this.cache.order)
      .set('dir', this.cache.dir);

    if (isFilter) {
      if (pageInfo.name) {
        params = params.append('name', pageInfo.name);
      }
      if (pageInfo.tags) {
        params.append('tags', JSON.stringify(pageInfo.tags));
      }
      this.paginationCache[offset] = false;
    }

    if (this.paginationCache[pageInfo.offset]) {
      return of(this.cache);
    }

    return this.http.get(environment.agentPoliciesUrl, { params })
      .map(
        (resp: any) => {
          this.paginationCache[pageInfo.offset] = true;
          // This is the position to insert the new data
          const start = resp.offset;
          const newData = [...this.cache.data];
          // TODO figure out what field name for object data in response...
          newData.splice(start, resp.limit, ...resp.data);
          this.cache = {
            ...this.cache,
            offset: Math.floor(resp.offset / resp.limit),
            total: resp.total,
            data: newData,
          };
          if (pageInfo.name) this.cache.name = pageInfo.name;
          if (pageInfo.tags) this.cache.tags = pageInfo.tags;
          return this.cache;
        },
      )
      .catch(
        err => {
          this.notificationsService.error('Failed to get Agent Policies',
            `Error: ${ err.status } - ${ err.statusText }`);
          return Observable.throwError(err);
        },
      );
  }

  getAvailableBackends() {
    // return this.http.get(environment.agentsBackendUrl)
    //   .map((resp: any) => {
    //     return resp.backend;
    //   }).catch(err => {
    //       this.notificationsService.error('Failed to get Available Backends',
    //         `Error: ${ err.status } - ${ err.statusText }`);
    //       return Observable.throwError(err);
    //     });
    // TODO uncomment mock above
    return new Observable(subscriber => {
      // TODO continue this format in future
      const resp = {
        data: [
          {
            'backend': 'pktvisor',
            'description': 'pktvisor observability agent from pktvisor.dev',
            // todo I could use some meta like this
            // 'config': ['taps', 'input', 'handlers'],
          },
        ],
      };
      subscriber.next(resp);
    });
  }

  // todo from this point on I have to assume pktvisor hardcoded steps
  // tap -> which will have a predefined input
  // fill input config form, will be dynamic to some extent
  // from there on, select handlers
  // ${backend}/${config[i]}/ // pktvisor/[taps,inputs,handlers]
  getBackendConfig(route: string[]) {
    const final = route.join('/');
    // return this.http.get(`${environment.agentsBackendUrl}/${final})
    //   .map((resp: any) => {
    //     return resp.backend;
    //   }).catch(
    //     err => {
    //       this.notificationsService.error('Failed to get Available Backends',
    //         `Error: ${ err.status } - ${ err.statusText }`);
    //       return Observable.throwError(err);
    //     },
    //   );
    // TODO remove mock and uncomment http request
    // if (final === 'pktvisor/taps') {
    //   return this.http.get(`${environment.agentsBackendUrl}/${final}`)
    //   .map((response: any) => {
    //     return response.backend;
    //   }).catch(
    //     err => {
    //       this.notificationsService.error('Failed to get Available Backends',
    //         `Error: ${ err.status } - ${ err.statusText }`);
    //       return Observable.throwError(err);
    //     },
    //   );
    // }

    let resp;
    switch (final) {
      case 'pktvisor/taps':
        resp = {
          data: [
            {
              'name': 'ethernet',
              'input_type': 'pcap',
              'config_predefined': {
                'iface': 'eth0',
              },
              'agents': {
                'total': 1,
              },
            },
          ],
        };
        break;
      case 'pktvisor/inputs':
        resp = {
          data: {
            'pcap': {
              'version': '1.0',
              'config': {
                'iface': {
                  'required': true,
                  'type': 'string',
                  'input': 'text',
                  'label': 'Interface',
                  'name': 'iface',
                  'description': 'The ethernet interface to capture on',
                },
                'bpf': {
                  'required': false,
                  'type': 'string',
                  'input': 'text',
                  'label': 'Filter Expression',
                  'name': 'bpf',
                  'description': 'tcpdump compatible filter expression for limiting the traffic examined (with BPF). Example: "port 53"',
                },
                'host_spec': {
                  'required': false,
                  'type': 'string',
                  'input': 'text',
                  'label': 'Host Specification',
                  'name': 'host_spec',
                  'description': 'Subnets (comma separated) to consider this HOST, in CIDR form. Example: "10.0.1.0/24,10.0.2.1/32,2001:db8::/64"',
                },
                'pcap_source': {
                  'required': false,
                  'type': 'string',
                  'input': 'text',
                  'label': 'pcap Engine',
                  'name': 'pcap_source',
                  'description': 'pcap backend engine to use. Defaults to best for platform.',
                },
              },
            },
            'dnstap': {
              'version': '1.0',
              'config': {
                'type': {
                  'type': 'string',
                  'input': 'select',
                  'label': 'Type',
                  'name': 'type',
                  'props': {
                    'options': {
                      'AUTH_QUERY': 'AUTH_QUERY',
                      'AUTH_RESPONSE': 'AUTH_RESPONSE',
                      'RESOLVER_QUERY': 'RESOLVER_QUERY',
                      'RESOLVER_RESPONSE': 'RESOLVER_RESPONSE',
                      'TOOL_QUERY': 'TOOL_QUERY',
                      'TOOL_RESPONSE': 'TOOL_RESPONSE',
                    },
                  },
                  'required': true,
                  'description': 'AUTH_QUERY, AUTH_RESPONSE, RESOLVER_QUERY, RESOLVER_RESPONSE, ..., TOOL_QUERY, TOOL_RESPONSE',
                },
                'socket_family': {
                  'type': 'string',
                  'input': 'select',
                  'label': 'Socket Family',
                  'name': 'socket_family',
                  'props': {
                    'options': {
                      'INET': 'INET',
                      'INET_6': 'INET_6',
                    },
                  },
                  'required': true,
                  'description': 'INET, INET6',
                },
                'socket_protocol': {
                  'type': 'string',
                  'input': 'select',
                  'label': 'Socket Protocol',
                  'name': 'socket_protocol',
                  'props': {
                    'options': {
                      'TCP': 'TCP',
                      'UDP': 'UDP',
                    },
                  },
                  'required': true,
                  'description': 'UDP, TCP',
                },
                'query_address': {
                  'type': 'string',
                  'input': 'text',
                  'label': 'Query Address',
                  'name': 'query_address',
                  'required': false,
                  'description': '',
                },
                'query_port': {
                  'type': 'string',
                  'input': 'text',
                  'label': 'Query Port',
                  'name': 'query_port',
                  'required': false,
                  'description': '',
                },
                'response_address': {
                  'type': 'string',
                  'input': 'text',
                  'label': 'Response Address',
                  'name': 'response_address',
                  'required': false,
                  'description': '',
                },
              },
            },
          },
        };
        break;
      case 'pktvisor/handlers':
        resp = {
          'data': {
            'dns': {
              'version': '1.0',
              'config': {
                'filter_exclude_noerror': {
                  'label': 'Filter: Exclude NOERROR',
                  'name': 'filter_exclude_noerror',
                  'type': 'boolean',
                  'input': 'checkbox',
                  'description': 'Filter out all NOERROR responses',
                },
                'filter_only_rcode': {
                  'label': 'Filter: Include Only RCode',
                  'name': 'filter_only_rcode',
                  'type': 'number',
                  'input': 'select',
                  'description': 'Filter out any queries which are not the given RCODE',
                  'props': {
                    'options': {
                      '0': '0 - NoError No Error [RFC1035]',
                      '1': '1 - FormErr Format Error [RFC1035]',
                      '2': '2 - ServFail Server Failure [RFC1035]',
                      '3': '3 - NXDomain Non Existent Domain [RFC1035]',
                      '4': '4 - NotImp Not Implemented [RFC1035]',
                      '5': '5 - Refused Query Refused [RFC1035]',
                      '6': '6 - YXDomain Name Exists when it should not	[RFC2136][RFC6672]',
                      '7': '7 - YXRRSet RR Set Exists when it should not	[RFC2136]',
                      '8': '8 - NXRRSet RR Set that should exist does not	[RFC2136]',
                      '9': '9 - NotAuth Server Not Authoritative for zone	[RFC2136] | NotAuth	Not Authorized	[RFC8945]',
                      '10': '10 - NotZone Name not contained in zone	[RFC2136]',
                      '11': '11 - DSOTYPENI DSO TYPE Not Implemented	[RFC8490]',
                      '16': '16 - BADVERS Bad OPT Version [RFC6891] | BADSIG	TSIG Signature Failure	[RFC8945]',
                      '17': '17 - BADKEY Key not recognized [RFC8945]',
                      '18': '18 - BADTIME Signature out of time window	[RFC8945]',
                      '19': '19 - BADMODE Bad TKEY Mode [RFC2930]',
                      '20': '20 - BADNAME Duplicate key name [RFC2930]',
                      '21': '21 - BADALG Algorithm not supported	[RFC2930]',
                      '22': '22 - BADTRUNC Bad Truncation [RFC8945]',
                      '23': '23 - BADCOOKIE Bad missing Server Cookie',
                    },
                  },
                },
                'filter_only_qname_suffix': {
                  'label': 'Filter: Include Only QName With Suffix',
                  'name': 'filter_only_qname_suffix',
                  'type': 'text',
                  'input': 'text',
                  'props': { 'pattern': '(\w+);' },
                  'description': 'Filter out any queries whose QName does not end in a suffix on the list',
                },
              },
            },
            'pcap': {
              'version': '1.0',
              'config': {},
            },
            'net': {
              'version': '1.0',
              'config': {},
            },
          },
        };
        break;
      default:
        resp = 'error';
    }
    return new Observable(subscriber => {
      subscriber.next(resp);
    });
  }

}
