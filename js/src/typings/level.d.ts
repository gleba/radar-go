// Type definitions for abstract-levelup

declare interface AbstractOptions {
  readonly [k: string]: any;
}

declare type ErrorCallback = (err: Error | undefined) => void;
declare type ErrorValueCallback<V> = (err: Error | undefined, value: V) => void;
declare type ErrorKeyValueCallback<K, V> = (err: Error | undefined, key: K, value: V) => void;

declare interface AbstractOpenOptions extends AbstractOptions {
  createIfMissing?: boolean;
  errorIfExists?: boolean;
}

declare interface AbstractGetOptions extends AbstractOptions {
  asBuffer?: boolean;
}

declare interface AbstractLevel<K = any, V = any> extends AbstractOptions {
  open(cb: ErrorCallback): void;
  open(options: AbstractOpenOptions, cb: ErrorCallback): void;

  close(cb: ErrorCallback): void;

  get(key: K, cb?: ErrorValueCallback<V>): void;
  get(key: K, options: AbstractGetOptions, cb?: ErrorValueCallback<V>): void;

  put(key: K, value: V, cb?: ErrorCallback): void;
  put(key: K, value: V, options: AbstractOptions, cb?: ErrorCallback): void;

  del(key: K, cb: ErrorCallback): void;
  del(key: K, options: AbstractOptions, cb: ErrorCallback): void;

  batch(): AbstractChainedBatch<K, V>;
  batch(array: any[], cb: ErrorCallback): AbstractChainedBatch<K, V>;
  batch(
    array: ReadonlyArray<AbstractBatch<K, V>>,
    options: AbstractOptions,
    cb: ErrorCallback,
  ): AbstractChainedBatch<K, V>;

  iterator(options?: AbstractIteratorOptions<K>): AbstractIterator<K, V>;
}

declare interface AbstractLevelDOWNConstructor {
  // tslint:disable-next-line no-unnecessary-generics
  new <K = any, V = any>(location: string): AbstractLevel<K, V>;
  // tslint:disable-next-line no-unnecessary-generics
  <K = any, V = any>(location: string): AbstractLevel<K, V>;
}

declare interface AbstractIteratorOptions<K = any> extends AbstractOptions {
  gt?: K;
  gte?: K;
  lt?: K;
  lte?: K;
  reverse?: boolean;
  limit?: number;
  keys?: boolean;
  values?: boolean;
  keyAsBuffer?: boolean;
  valueAsBuffer?: boolean;
}

declare type AbstractBatch<K = any, V = any> = PutBatch<K, V> | DelBatch<K, V>;

declare interface PutBatch<K = any, V = any> {
  readonly type: 'put';
  readonly key: K;
  readonly value: V;
}

declare interface DelBatch<K = any, V = any> {
  readonly type: 'del';
  readonly key: K;
}

declare interface AbstractChainedBatch<K = any, V = any> extends AbstractOptions {
  put: (key: K, value: V) => this;
  del: (key: K) => this;
  clear: () => this;
  write(cb: ErrorCallback): any;
  write(options: any, cb: ErrorCallback): any;
}

declare interface AbstractChainedBatchConstructor {
  // tslint:disable-next-line no-unnecessary-generics
  new <K = any, V = any>(db: any): AbstractChainedBatch<K, V>;
  // tslint:disable-next-line no-unnecessary-generics
  <K = any, V = any>(db: any): AbstractChainedBatch<K, V>;
}

declare interface AbstractIterator<K, V> extends AbstractOptions {
  //db: AbstractLevelDOWN<K, V>;
  next(cb: ErrorKeyValueCallback<K, V>): this;
  end(cb: ErrorCallback): void;
}

declare interface AbstractIteratorConstructor {
  // tslint:disable-next-line no-unnecessary-generics
  new <K = any, V = any>(db: any): AbstractIterator<K, V>;
  // tslint:disable-next-line no-unnecessary-generics
  <K = any, V = any>(db: any): AbstractIterator<K, V>;
}

