// Esta función toma el valor del hash y lo utiliza para invocar la función correspondiente
export function handleHashChange(): string {
  const hash = window.location.hash.slice(1); // Obtener el valor del hash sin el '#'
  
  // Verificar si hay un valor en el hash
  if (hash) {
    // Invocar la función con el nombre del hash
    if (typeof window[hash] === 'function') {
      window[hash](); // Invocar la función
    } else {
      console.error(`The '${hash}' function is not definied.`);
    }
    return hash
  }
}

// add this code:
// window.addEventListener('DOMContentLoaded', handleHashChange);
// window.addEventListener('hashchange', handleHashChange);
// let route=handleHashChange();