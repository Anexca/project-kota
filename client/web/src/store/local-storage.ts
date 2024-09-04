import { create } from "zustand";
import { persist } from "zustand/middleware";

// Define your state interface
interface StoreState {
  items: string[]; // Array of strings
  addItem: (item: string) => void; // Add item to the array
  removeItem: (item: string) => void; // Remove item from the array
  clearItems: () => void; // Clear all items
}

// Create the store with persistence
const useLocalStorageStore = create(
  persist<StoreState>(
    (set) => ({
      items: [], // Initialize with an empty array
      addItem: (item: string) =>
        set((state) => ({ items: [...state.items, item] })), // Add item to the array
      removeItem: (item: string) =>
        set((state) => ({
          items: state.items.filter((i) => i !== item), // Remove item from the array
        })),
      clearItems: () => set(() => ({ items: [] })), // Clear the array
    }),
    {
      name: "attempted-questions", // Unique name for localStorage
      getStorage: () => localStorage, // Specify that we're using localStorage
    }
  )
);

export default useLocalStorageStore;
