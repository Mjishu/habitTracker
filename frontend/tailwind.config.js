/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./app/**/*.{js,jsx,ts,tsx}"],
  presets: [require("nativewind/preset")],
  theme: {
    extend: {
      colors: {
        primary: "#0e000e",
        secondary: "#51227a",
        accent: "#e0d964",
        black: {
          100: "#1c191e",
        },
      },
    },
  },
  plugins: [],
};
