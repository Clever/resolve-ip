import { Logger } from "kayvee";

type Callback<R> = (err: Error, result: R) => void;
type ArrayInner<R> = R extends (infer T)[] ? T : never;

interface RetryPolicy {
  backoffs(): number[];
  retry(requestOptions: {method: string}, err: Error, res: {statusCode: number}): boolean;
}

interface RequestOptions {
  timeout?: number;
  retryPolicy?: RetryPolicy;
}

interface IterResult<R> {
  map<T>(f: (r: R) => T, cb?: Callback<T[]>): Promise<T[]>;
  toArray(cb?: Callback<R[]>): Promise<R[]>;
  forEach(f: (r: R) => void, cb?: Callback<void>): Promise<void>;
  forEachAsync(f: (r: R) => void, cb?: Callback<void>): Promise<void>;
}

interface CircuitOptions {
  forceClosed?: boolean;
  maxConcurrentRequests?: number;
  requestVolumeThreshold?: number;
  sleepWindow?: number;
  errorPercentThreshold?: number;
}

interface GenericOptions {
  timeout?: number;
  keepalive?: boolean;
  retryPolicy?: RetryPolicy;
  logger?: Logger;
  circuit?: CircuitOptions;
  serviceName?: string;
}

interface DiscoveryOptions {
  discovery: true;
  address?: undefined;
}

interface AddressOptions {
  discovery?: false;
  address: string;
}

type ResolveIPOptions = (DiscoveryOptions | AddressOptions) & GenericOptions;

import models = ResolveIP.Models

declare class ResolveIP {
  constructor(options: ResolveIPOptions);

  close(): void;
  
  healthCheck(options?: RequestOptions, cb?: Callback<void>): Promise<void>
  
  locationForIP(ip: string, options?: RequestOptions, cb?: Callback<models.IP>): Promise<models.IP>
  
}

declare namespace ResolveIP {
  const RetryPolicies: {
    Single: RetryPolicy;
    Exponential: RetryPolicy;
    None: RetryPolicy;
  }

  const DefaultCircuitOptions: CircuitOptions;

  namespace Errors {
    interface ErrorBody {
      message: string;
      [key: string]: any;
    }

    
    class BadRequest {
  message?: string;

  constructor(body: ErrorBody);
}
    
    class InternalError {
  message?: string;

  constructor(body: ErrorBody);
}
    
    class NotFound {
  message?: string;

  constructor(body: ErrorBody);
}
    
  }

  namespace Models {
    
    type IP = {
  lat: number;
  lon: number;
};
    
  }
}

export = ResolveIP;
