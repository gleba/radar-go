// Copyright (c) Gleb Panteleev. All rights reserved. Licensed under the MIT license.

/**
 * Набор хуков для Preact
 * @remarks
 * @packageDocumentation
 */

import { A, IAtom } from 'alak'
import { useCallback, useEffect, useState } from 'preact/hooks'

export const alive = (v) => (v !== undefined && v !== null) as boolean

export function useAtom<T>(atom: IAtom<T>): T {
  const [state, fx] = useState(atom.value)

  // @ts-ignore
  useEffect(() => (atom.up(fx), () => atom.down(fx)))
  return state
}

export function useComputeAtom<T, U>(
  atom: IAtom<T>,
  computeFn: (v: T) => U
): [U] {
  let lastValue = atom.value
  let value = computeFn(lastValue)
  const [state, mutate] = useState(value)
  let mutateFx = (v) => {
    if (lastValue !== v) {
      lastValue = v
      mutate(computeFn(v))
    }
  }
  useEffect(() => {
    atom.up(mutateFx)
    return () => atom.down(mutateFx)
  }, [atom])
  return [state]
}

export function useAtomFx<T>(atom: IAtom<T>, effectFn: (v: T) => void): [T] {
  let lastValue = atom.value
  const [state, mutate] = alive(lastValue)
    ? useState(lastValue)
    : useState(undefined)
  let mutateFx = (v) => {
    if (lastValue !== v) {
      lastValue = v
      effectFn(v)
      mutate(v)
    }
  }
  useEffect(() => {
    atom.up(mutateFx)
    return () => atom.down(mutateFx)
  }, [atom])
  return [state]
}

export function useASyncAtom<T, U>(
  atom: IAtom<T>,
  mixin?: (v: T) => U
): [U, Boolean] {
  const [state, mutate] = alive(atom.value)
    ? useState(atom.value)
    : useState(undefined)
  // let busy
  // if (atom.isAsync) {
  const [now, change] = useState(false)
  // busy = { now, change }
  // }
  useEffect(() => {
    const mutator = (v) => state !== v && mutate(v)
    atom.up(mutator)
    atom.onAwait(change)
    return () => {
      atom.down(mutator)
      atom.onAwait(change)
    }
  }, [atom])
  return [state, now]
}

const asEventHandler = (e, value) => {
  // const [e, value] = a
  if (value != undefined) return value
  if (e.target) {
    if ('value' in e.target) return e.target.value
    if ('checked' in e.target) return e.target.checked
  }
  return ''
}

export function useInputAtom<T>(
  atom: IAtom<T>,
  effectFn?: (v: T) => void
): [T, any] {
  let lastValue = atom.value
  const [state, mutate] = alive(lastValue)
    ? useState(lastValue)
    : useState(undefined)
  const mutateFx = (v) => {
    if (lastValue !== v) {
      if (atom.value != v) atom(v)
      lastValue = v
      if (effectFn) effectFn(v)
      mutate(v)
    }
  }
  // @ts-ignore
  const eventHandler = (...a) => mutateFx(asEventHandler(...a))
  useEffect(() => {
    atom.up(mutateFx)
    return () => atom.down(mutateFx)
  }, [atom])
  return [state, eventHandler]
}

export function useOnAtom<T>(
  atom: IAtom<T>,
  listingFn: (v: T) => void,
  ...diff: any[]
): void {
  useEffect(() => {
    atom.up(listingFn)
    return () => atom.down(listingFn)
  }, [atom, ...diff])
}
