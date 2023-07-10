function qs(arr: number[], low: number, high: number): void {
  // base case
  if (low >= high) {
    return;
  }

  const pivotIdx = partition(arr, low, high);

  qs(arr, low, pivotIdx - 1);
  qs(arr, pivotIdx+1, high);
}

function partition(arr: number[], low: number, high: number): number {
  // nothing to fancy
  const pivot = arr[high];

  // idx start from low - 1
  let idx = low - 1;

  for (let i = low; i < high; i++){
    if (arr[i] <= pivot) {
      // start the idx
      idx++;

      // swap the value
      const tmp = arr[i];
      arr[i] = arr[idx];
      arr[idx] = tmp;
    }
  }

   // point to pivot idx
  idx++;

   // swap the value
  arr[high] = arr[idx];
  arr[idx] = pivot;

  return idx;
}

export default function quick_sort(arr: number[]): void {
  qs(arr, 0, arr.length - 1);
}