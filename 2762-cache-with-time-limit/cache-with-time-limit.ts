type MapEntry = { value: number, timeout: NodeJS.Timeout };

class TimeLimitedCache {
  cache = new Map<number, MapEntry>();

  set(key: number, value: number, duration: number) {
    const valueInCache = this.cache.get(key);
    if (valueInCache) {
      clearTimeout(valueInCache.timeout);
    }
    const timeout = setTimeout(() => this.cache.delete(key), duration);
    this.cache.set(key, { value, timeout });
    return Boolean(valueInCache);
  }

  get(key: number) {
    return this.cache.has(key) ? this.cache.get(key).value : -1;
  }

  count() {
    return this.cache.size;
  }
};