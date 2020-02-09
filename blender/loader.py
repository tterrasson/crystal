import os
import bpy
import math


INPUT_PATH = '/home/...../crystal/explore'
OUTPUT_PATH = '/home/...../crystal/output'
RESOLUTION_X = 1920
RESOLTUION_Y = 1080
ITERATION = 8


# Delete objects
def delete_all_mesh():
    for obj in bpy.context.scene.objects:
        if obj.type == 'MESH':
            obj.select_set(True)
        else:
            obj.select_set(False)

    # Call the operator only once
    bpy.ops.object.delete()


def render(idx):
    filepath = os.path.join(INPUT_PATH, f'output-{idx}.obj')
    imported_object = bpy.ops.import_scene.obj(filepath=filepath)
    obj_object = bpy.context.selected_objects[0]
    bpy.ops.object.origin_set(type='GEOMETRY_ORIGIN', center='MEDIAN')
    print(f'Imported name: {obj_object.name}')

    # Zoom to object
    for area in bpy.context.screen.areas:
        if area.type == 'VIEW_3D':
            for region in area.regions:
                if region.type == 'WINDOW':
                    override = {'area': area, 'region': region}
                    bpy.ops.view3d.camera_to_view_selected(override)

    # Render frame
    bpy.context.scene.render.filepath = os.path.join(OUTPUT_PATH, f'output-{idx}.png')
    bpy.ops.render.render(write_still=True)


for i in range(ITERATION):
    delete_all_mesh()
    idx = f"{i}".zfill(6)
    print(f"Rendering {idx}...")
    render(idx)