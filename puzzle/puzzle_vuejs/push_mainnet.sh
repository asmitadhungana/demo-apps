echo "I assumed you already built (npm run build) before this pushing"
# yarn build
firebase deploy --only hosting:puzzlemainnet
