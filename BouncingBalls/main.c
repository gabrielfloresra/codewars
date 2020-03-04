#include <stdio.h>


int checkParameters(double heightChildPlaying, double ballBounces, double windowHeight);
int getManyTimesSeeeTheBall(double heightChildPlaying, double ballBounces, double windowHeight);

int main(int argc, char const *argv[])
{
    double heightChildPlaying = 3000000000000000000;
    double ballBounces = 0.99999999;
    double windowHeight = 0.1;
    // double heightChildPlaying = 30;
    // double ballBounces = 0.66;
    // double windowHeight = 1.5;

    int manyTimesSeeTheBall = getManyTimesSeeeTheBall(heightChildPlaying, ballBounces, windowHeight);

    printf("she can see %d times the ball", manyTimesSeeTheBall);

    return 0;
}

int checkParameters(double heightChildPlaying, double ballBounces, double windowHeight)
{
    int condition1 = 1;
    int condition2 = 1;
    int condition3 = 1;

    if (heightChildPlaying <= (double)0)
    {
        condition1 = 0;
        printf("\tcondition 1 not fulfilled\n");
    }
    if (ballBounces <= (double)0 || ballBounces > (double)1)
    {
        condition2 = 0;
        printf("\tcondition 2 not fulfilled\n");
    }
    if (windowHeight <= (double)0 || windowHeight >= heightChildPlaying)
    {
        condition3 = 0;
        printf("\tcondition 3 not fulfilled\n");
    }

    return condition1 && condition2 && condition3;
}

int getManyTimesSeeeTheBall(double heightChildPlaying, double ballBounces, double windowHeight)
{
    int manyTimesSeeTheBall = -1;

    if (checkParameters(heightChildPlaying, ballBounces, windowHeight))
    {
        double heightReboundingBall = heightChildPlaying;
        manyTimesSeeTheBall = 0;
        while (heightReboundingBall > windowHeight)
        {
            // printf("cae desde: %f", heightReboundingBall);
            manyTimesSeeTheBall += 2;
            heightReboundingBall *= ballBounces;
            // printf("\nsube hasta: %.3f\n", heightReboundingBall);
        }
        if (heightReboundingBall <= windowHeight)
        {
            --manyTimesSeeTheBall;
        }
    }

    return manyTimesSeeTheBall;
}