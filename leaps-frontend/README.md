# LEAPS Frontend - Web Portal

Leadership Preparatory Schools Digital Operating System Frontend

## Setup Instructions

### Prerequisites
- Modern web browser (Chrome, Firefox, Safari, Edge)
- Node.js 16+ (optional, for local server)
- Backend API running on `http://localhost:8080`

### Quick Start

1. **Navigate to frontend directory**
```bash
cd leaps-frontend
```

2. **Start a local server** (optional)
```bash
# Using Python 3
python -m http.server 3000

# Using Node.js
npx http-server -p 3000
```

3. **Open in browser**
```
http://localhost:3000
```

## Pages

### Public Pages
- `index.html` - Landing page
- `login.html` - User login

### Protected Pages
- `dashboard.html` - Main dashboard (role-based)

## Features

### Authentication
- Email/password login
- JWT token storage
- Automatic logout on token expiry
- Session persistence

### Role-Based Dashboards
- **Owner** - Full system control
- **Principal** - Secondary academic management
- **Head Teacher** - Primary/Nursery management
- **Teacher** - Score entry and class management
- **Bursar** - Finance management
- **Student** - Results and quizzes
- **Parent** - Child monitoring

### Modules

#### Students
- View student directory
- Add new students
- Update student information
- Track student progress

#### Admissions
- View applications
- Schedule interviews
- Approve/reject admissions
- Generate admission letters

#### Academic Results
- Submit scores
- View computed results
- Publish results
- Lock results for finality

#### Quizzes
- Create quizzes
- Manage questions
- Monitor quiz attempts
- View results

#### Finance
- Manage fees
- Record payments
- Track debts
- Generate receipts

#### Reports
- Generate report cards
- Download documents
- Print reports
- Email reports

## Project Structure

```
leaps-frontend/
├── index.html                   # Landing page
├── login.html                   # Login page
├── dashboard.html               # Main dashboard
├── assets/
│   ├── css/
│   │   ├── base.css            # Base styles
│   │   ├── auth.css            # Auth page styles
│   │   ├── dashboard.css       # Dashboard styles
│   │   └── animations.css      # Animations
│   ├── js/
│   │   ├── api.js              # API client
│   │   ├── auth.js             # Auth logic
│   │   ├── router.js           # Page router
│   │   └── dashboard.js        # Dashboard logic
│   └── images/
│       ├── logo.png            # School logo
│       └── avatar.png          # User avatar
└── pages/
    ├── owner/                  # Owner pages
    ├── principal/              # Principal pages
    ├── teacher/                # Teacher pages
    ├── student/                # Student pages
    └── parent/                 # Parent pages
```

## API Integration

The frontend communicates with the backend via REST API:

```javascript
// Example API call
const response = await AuthAPI.login(email, password);
if (response.success) {
  localStorage.setItem("token", response.data.token);
}
```

## Styling

### Color Scheme
- Primary: `#1e3a8a` (Dark Blue)
- Secondary: `#fbbf24` (Amber)
- Success: `#10b981` (Green)
- Danger: `#ef4444` (Red)
- Background: `#f8fafc` (Light Gray)

### Design System
- Glassmorphism effects
- Soft shadows
- Card-based layout
- Smooth animations
- Mobile-responsive

## Browser Support

- Chrome 90+
- Firefox 88+
- Safari 14+
- Edge 90+

## Performance

- Lazy loading of pages
- Minimal dependencies
- Optimized CSS
- Efficient JavaScript
- Responsive images

## Security

- JWT token authentication
- Secure token storage
- CORS enabled
- Input validation
- XSS protection

## Development

### Adding New Pages

1. Create HTML file in `pages/` directory
2. Add route in `router.js`
3. Create CSS file in `assets/css/`
4. Add navigation link in `dashboard.html`

### Customization

Edit CSS variables in `base.css`:
```css
:root {
  --primary: #1e3a8a;
  --secondary: #fbbf24;
  /* ... */
}
```

## License

Proprietary - Leadership Preparatory Schools
