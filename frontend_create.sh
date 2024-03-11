#!/bin/bash

# Instalar Vite y Svelte
rm -rf ./node_modules
npm init -y
npm install -D create-vite
npx create-vite ./ --template svelte-ts

# Instalar Tailwind CSS
npm install -D tailwindcss postcss autoprefixer
npx tailwindcss init -p

# Configurar PostCSS
cat <<EOL > postcss.config.js
module.exports = {
  plugins: {
    tailwindcss: {},
    autoprefixer: {},
  },
};
EOL

# Instalar Vite plugins
npm install -D vite-plugin-svelte
npm install -D vite-plugin-fonts

# Configurar Vite para Svelte y TypeScript
cat <<EOL > vite.config.js
import { defineConfig } from 'vite';
import Svelte from 'vite-plugin-svelte';
import ViteFonts from 'vite-plugin-fonts';

export default defineConfig({
  plugins: [Svelte(), ViteFonts()],
  build: {
    target: 'esnext',
  },
});
EOL

# Crear el archivo de estilos de Tailwind
mkdir src/styles
cat <<EOL > src/styles/tailwind.css
@import 'tailwindcss/base';
@import 'tailwindcss/components';
@import 'tailwindcss/utilities';
EOL

# Instalar Flowbite
npm install flowbite

# Ajustar el punto de entrada principal
echo "Asegúrate de ajustar el punto de entrada principal en 'src/main.ts'."

# Importar Flowbite en el punto de entrada principal
echo "Importa Flowbite en 'src/main.ts'"
echo "import 'flowbite/dist/flowbite.css';" >> src/main.ts

# Mostrar mensaje final
echo "Listo! Tu proyecto Vite con Svelte, Tailwind CSS, TypeScript y Flowbite está configurado."
