#include <stdio.h>
#include <stdlib.h>

struct Node
{
    int val;
    struct Node *next;
};

struct Node* createNewNode(int val, struct Node *next) {
    struct Node *el = (struct Node*)malloc(sizeof(struct Node));
    el->val = val;
    el->next = next;
    return el;
}

struct Node *curr;

struct Node* reverseLL(struct Node *head)  {
    struct Node *prev = NULL;
    curr = head;


    while (curr != NULL)
    {
        struct Node *next = curr->next;
        curr->next = prev;
        prev = curr;
        curr = next;
    }
    
    return prev;
};

void insertNthEl(int nth, struct Node *head, struct Node *newEl){
    curr = head;
    int ctr = 1;
    while (curr != NULL) 
    {
        if (ctr == nth - 1) {
            newEl->next = curr->next;
            curr->next = newEl;
            curr = newEl;
            break;
        }            
        curr = curr->next;    
        ctr++;
    }
}

void printLL(struct Node *head) {
    curr = head;
    while (curr != NULL)
    {
        printf("\nValue is: %d\n",curr->val);
        curr = curr->next;
    }
    
}

int main() {
    struct Node 
            *l11 = NULL, 
            *l12 = NULL,
            *l13 = NULL,
            *l14 = NULL,
            *l15 = NULL,
            *newEl = NULL;

    l15 = createNewNode(15, NULL);
    l14 = createNewNode(14, l15);
    l13 = createNewNode(14, l14);
    l12 = createNewNode(14, l13);
    l11 = createNewNode(14, l12);

    newEl = createNewNode(9999, NULL);
    printLL(l11);
    insertNthEl(3, l11,newEl);
    printf("=========================");
    printLL(l11);

    // struct Node *res = reverseLL(l11);
    // printf("=========================");
    // printLL(res);
    return 0;
}

