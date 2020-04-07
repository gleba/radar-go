const si = {
  Period1: 0,
  Period7: 1,
  Period30: 2,
  Period60: 3,
  Period45: 4,
  Period90: 5,
}

export const tgl = {
  prices: {
    median: 0,
    mean: 1,
    harmonic: 2,
    // PriceGeometric: 3,
    variance: 4,
  },
  volumes: {
    median: 5,
    mean: 6,
    // VolumeVariance: 7,
    harmonic: 8,
    // VolumeGeometric: 9
  },

  volatile: {
    mean: 10,
    median: 11,
    // VolatilityGeometric: 12,
    harmonic: 13,
  },
  percentile: {
    nearestRank30: 14,
    Rank60: 15,
    Rank90: 16,
    30: 17,
    60: 18,
    90: 19,
  },
}
