import { performance } from 'perf_hooks';

export default function benchmark(name, fn, args) {
  let t = performance.now()
  fn(...args)
  console.log(`function ${name} took ${(performance.now() - t)} milliseconds`)
}