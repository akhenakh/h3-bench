// gcc -o c main.c -lh3

#include <h3/h3api.h>
#include <inttypes.h>
#include <stdio.h>

int main(int argc, char *argv[])
{
    for (int i = 0; i < 1000000; i++)
    {
        LatLng location;
        location.lat = degsToRads(48.8);
        location.lng = degsToRads(2.2);

        H3Index indexed;
        if (latLngToCell(&location, 15, &indexed) != E_SUCCESS)
        {
            printf("Failed\n");
            return 1;
        }
        cellToLatLng(indexed, &location);
    }
}