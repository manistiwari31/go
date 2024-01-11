/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
    if root == nil{
        return 0
    }
    maxDiff := 0
    maxSoFar := root.Val
    minSoFar := root.Val
    meAndMyAncestor(root, maxSoFar,minSoFar, &maxDiff)
    return maxDiff
}

func meAndMyAncestor(root *TreeNode, maxSoFar int, minSoFar int, maxDiff *int) {
    if root == nil{
        return 
    }

    if root.Val > maxSoFar{
        maxSoFar = root.Val
    }else if root.Val < minSoFar{
        minSoFar = root.Val
    }
    
    
    if (maxSoFar - root.Val) > *maxDiff{
        *maxDiff = maxSoFar - root.Val
    }
    if (root.Val - minSoFar) > *maxDiff{
        *maxDiff = root.Val - minSoFar
    }
    //fmt.Println(maxSoFar, *maxDiff)
    meAndMyAncestor(root.Left, maxSoFar, minSoFar, maxDiff)
    meAndMyAncestor(root.Right, maxSoFar, minSoFar, maxDiff)
    return


}

func abs (a int) int{

    if a < 0 {
        return -a
    }else {
        return a
    }

}