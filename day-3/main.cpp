#include <iostream>
#include <string>
#include <fstream>

#define MAP_HEIGHT 323
#define MAP_WIDTH 31

const char MAP_SPACE = '.';
const char MAP_TREE = '#';

typedef struct Vector2i Vector2i;
struct Vector2i {
    int x;
    int y;
};

void read_input(std::string map[MAP_HEIGHT])
{
    std::ifstream file ("input.txt");
    for (int i = 0; i < MAP_HEIGHT; ++i) {
        file >> map[i];
    }
    file.close();
}

bool move(
    std::string map[MAP_HEIGHT],
    Vector2i& current_position,
    Vector2i leap_size
) {
    for (int x = 0; x < leap_size.x; ++x) {
        current_position.x++;
        if (current_position.x >= 31) current_position.x = 0;
    }
    current_position.y += leap_size.y;
    if (current_position.y >= MAP_HEIGHT) {
        return false;
    }
    return map[current_position.y][current_position.x] == MAP_TREE;
}

int main()
{
    std::string base_map[MAP_HEIGHT];
    read_input(base_map);
    Vector2i current_position = {0, 0};
    bool has_tree;
    int total_trees = 0;
    while (current_position.y < MAP_HEIGHT) {
        has_tree = move(base_map, current_position, {1, 2});
        if (has_tree) total_trees++;
    }
    printf("%d\n", total_trees);
    return 0;
}
