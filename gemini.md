# Project: SocialSync

## Overview
SocialSync is a full-stack social media management tool for influencers and teams to create, schedule, and publish content across multiple social platforms (Facebook, Instagram, X/Twitter, YouTube, Mastodon). It supports workspace collaboration, post analytics, shared media libraries, and OAuth integration.

## Stack
- Frontend: Next.js, Tailwind
- Backend: Go (Golang), Gorilla Mux
- Database: Neon Postgres
- Media: Cloudinary
- Hosting: Vercel (frontend), Render (backend)

## Features
- ✅ JWT-based auth + Google OAuth
- ✅ Social account linking (FB, IG, Twitter, etc.)
- ✅ Post creation
- ✅ Image/video media upload
- 🚧 Draft management
- 🚧Workspace + roles (strategist, reviewer, publisher)
- 🚧 Analytics from Facebook/Instagram APIs
- 🚧 Shared calendar, comments, campaign planning

## Workflows

### 1. OAuth Connection
1. User initiates social connect
2. Redirect to platform login
3. Backend callback saves access/refresh tokens

### 2. Post Scheduling
1. User drafts post with message/media
2. Selects platform(s) and schedule time
3. Worker picks it up later and publishes via API

### 3. Analytics
1. Fetch insights from Facebook/Instagram Graph API
2. Store in `post_metrics` table
3. Render charts in dashboard using Recharts

### 4. Media Upload
1. User uploads to Cloudinary
2. Store resulting URL
3. Attach to post draft/published

### 5. Team Collaboration
1. Workspace creation
2. Role-based access
3. Review/comment/approve posts

## TODO
- [ ] Post insights dashboard
- [ ] External post syncing
- [ ] Content calendar
- [ ] Team feedback system
- [ ] Campaign-based planning
