class Solution {
public:
    bool judgeCircle(string moves) {
        
        int position[2] = {0, 0};

        for (char move : moves) {
            switch (move) {
                case 'U':
                    position[0]++;
                    break;
                case 'D':
                    position[0]--;
                    break;
                case 'L':
                    position[1]--;
                    break;
                case 'R':
                    position[1]++;
                    break;
            }
        }

        return position[0] == 0 && position[1] == 0;
    }
};
