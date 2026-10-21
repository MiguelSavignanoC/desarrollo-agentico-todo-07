import { test, expect } from '../fixtures';

test.describe('Tareas', () => {
  test('cambiar y quitar la fecha límite', async ({ page, tasks }) => {
    await tasks([{ title: 'Llamar al proveedor', dueDate: '2026-10-30' }]);
    await page.reload();
    await expect(page.locator('.task-due-date')).toHaveText(['Fecha límite: 30/10/2026']);

    // 1. Abre la edición: el campo trae la fecha actual
    await page.getByRole('link', { name: 'Editar tarea: Llamar al proveedor' }).click();
    await expect(page.getByLabel('Fecha límite')).toHaveValue('2026-10-30');

    // 2. Cambia la fecha y guarda
    await page.getByLabel('Fecha límite').fill('2026-11-15');
    await page.getByRole('button', { name: 'Guardar cambios' }).click();
    await expect(page.locator('.task-due-date')).toHaveText(['Fecha límite: 15/11/2026']);

    // 3. Vacía el campo y guarda: la tarea se queda sin fecha
    await page.getByRole('link', { name: 'Editar tarea: Llamar al proveedor' }).click();
    await page.getByLabel('Fecha límite').fill('');
    await page.getByRole('button', { name: 'Guardar cambios' }).click();
    await expect(page.getByRole('status')).toHaveText('Tarea actualizada.');
    await expect(page.locator('.task-due-date')).toHaveCount(0);

    // 4. Al recargar sigue sin fecha
    await page.reload();
    await expect(page.locator('.task-due-date')).toHaveCount(0);
  });
});
