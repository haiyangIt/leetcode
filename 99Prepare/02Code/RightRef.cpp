#include <iostream>
#include <vector>
#include <memory>
#include <algorithm>
using namespace std;

class Solution
{
public:
    vector<int> FindElement(vector<int> &array, int find);
    void Output(vector<int> &array);
};

vector<int> Solution::FindElement(vector<int> &array, int find)
{
    vector<int> result;
    if (array.size() == 0)
        return result;
    int left = 0;
    int right = array.size() - 1;
    int length = array.size();

    return result;
}

class MyBase
{
public:
    int a;
    int b;

protected:
    int c;
};

class MyDerived:MyBase{
public:
    MyDerived *p;
    void output(){
        cout <<p->c;
    }
};

union MyUnion {
    int A;
    char C[4];
};

void Solution::Output(vector<int> &array)
{
    for (auto first = array.begin(); first != array.end(); first++)
    {
        cout << *first << endl;
    }
}

int main()
{
    int x = 10;
    
    auto add_x = [=](int a) mutable { x *= 2; return a + x; };  // 复制捕捉x
    
    cout << add_x(10) << endl; // 输出 30
    cout << x <<endl;

    MyUnion u;
    u.A = 0x12345678;
    printf("%0xd", (int)u.C[0]);
    cout <<(int)u.C[0] <<endl;
    cout << "hello world!" << endl;
    vector<int> test;
    test.push_back(1);
    test.push_back(2);
    test.push_back(5);
    test.push_back(7);
    test.push_back(6);
    Solution s;
    s.Output(test);
    std::sort(test.begin(), test.end());
    s.Output(test);
    char c;
    Solution sa[2];
    cout << &(sa[0]) << endl;
    cout << &(sa[1]) << endl;
    cin >> c;
}
