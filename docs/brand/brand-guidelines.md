# NeuralMesh Brand Guidelines

Version 1.0 — May 2026

---

## 1. Name and Tagline

**Name:** NeuralMesh

- Always written as one word, capital N, capital M.
- Never: `neuralmesh`, `neural mesh`, `NEURALMESH`, `Neural-Mesh`.

**Tagline:** "Mine Intelligence, Not Waste."

- Full stop included.
- Use in marketing copy, not in UI or API documentation.

**Tone:** Technical, precise, accessible. Not crypto-speculative. Not hype-driven.

---

## 2. Color Palette

### Primary Colors

| Name | Hex | Usage |
|------|-----|-------|
| Space Black | `#0A0E1A` | Backgrounds, dark surfaces |
| Cyan Electric | `#00D4FF` | Primary accent, nodes, CTAs, links |
| Deep Violet | `#7B2FFF` | Secondary accent, center node, gradients |

### Secondary Colors

| Name | Hex | Usage |
|------|-----|-------|
| Off White | `#E8EAED` | Primary text on dark backgrounds |
| Slate Gray | `#4A5568` | Secondary text, labels, captions |
| Mid Gray | `#9CA3AF` | Muted text, placeholders |
| Neon Green | `#00FF88` | Success states, verified/reward indicators only |

### Usage Rules

- Never use Cyan Electric on white backgrounds (contrast ratio < 3:1).
- Neon Green is reserved for verified computation and reward states only.
- Dark surfaces use Space Black (`#0A0E1A`) or ≥ 90% equivalent.
- On light backgrounds, use the mono logo variant.

---

## 3. Typography

### Fonts

| Role | Font | Fallback | Source |
|------|------|----------|--------|
| Heading | Space Grotesk | Inter, sans-serif | [Google Fonts](https://fonts.google.com/specimen/Space+Grotesk) |
| Body | Inter | system-ui, sans-serif | [Google Fonts](https://fonts.google.com/specimen/Inter) |
| Monospace | JetBrains Mono | monospace | [Google Fonts](https://fonts.google.com/specimen/JetBrains+Mono) |

### Type Scale

| Level | Font | Size | Weight | Letter-spacing |
|-------|------|------|--------|----------------|
| H1 | Space Grotesk | 48px | 700 | -1px |
| H2 | Space Grotesk | 32px | 700 | -0.5px |
| H3 | Space Grotesk | 24px | 600 | 0 |
| Body | Inter | 16px | 400 | 0 |
| Caption | Inter | 12px | 400 | +1px |
| Label | Inter | 11px | 400 | +2.5px uppercase |
| Code | JetBrains Mono | 14px | 400 | 0 |

---

## 4. Logo

### Icon

The NeuralMesh icon is a **hexagonal mesh**: 7 nodes (1 center + 6 outer) connected
by spokes and a ring. The center node uses Deep Violet with a Cyan Electric core,
representing the intelligence hub of the network.

### Variants

| File | Background | Use case |
|------|------------|----------|
| `neuralmesh-logo-color.svg` | Space Black (`#0A0E1A`) | Default — website, presentations |
| `neuralmesh-logo-dark.svg` | Transparent | Overlay on dark or image backgrounds |
| `neuralmesh-logo-mono.svg` | White | Print, documents, light backgrounds |
| `favicon.svg` | Space Black | Browser tab, 32×32 and 192×192 |

### Minimum Size

- Horizontal lockup (icon + wordmark): minimum 200px wide
- Icon only: minimum 24px

### Clear Space

Maintain clear space of ≥ 1× the icon height on all sides.

### Prohibited Uses

- Do not recolor the logo.
- Do not stretch or distort proportions.
- Do not place the color variant on white or light backgrounds.
- Do not add drop shadows, outlines, or effects not in the source SVG.
- Do not use the tagline as part of the logo lockup.

---

## 5. Icon Only

For favicons, app icons, and avatar contexts, use the icon without the wordmark
(derived from `favicon.svg`). Export at:

- 32×32 px (browser favicon)
- 192×192 px (PWA / Android)
- 512×512 px (app store, high-res)

PNG exports at these sizes are generated from `favicon.svg` using Inkscape or similar:

```bash
inkscape favicon.svg --export-png=favicon-32.png --export-width=32
inkscape favicon.svg --export-png=favicon-192.png --export-width=192
inkscape favicon.svg --export-png=favicon-512.png --export-width=512
```

---

## 6. Voice and Tone

| Context | Tone |
|---------|------|
| Technical docs | Precise, factual, no marketing language |
| Marketing copy | Confident, clear, no hype or financial promises |
| Community | Accessible, inclusive, technically grounded |
| Error messages | Direct, helpful, no blame |

**Avoid:** "revolutionary", "disruptive", "to the moon", "passive income", "guaranteed returns".

**Use:** "verified", "distributed", "useful computation", "earn by contributing", "open protocol".

---

## 7. Assets Location

```text
assets/
└── brand/
    ├── neuralmesh-logo-color.svg    # Default (dark bg)
    ├── neuralmesh-logo-dark.svg     # White, transparent bg
    ├── neuralmesh-logo-mono.svg     # Black, white bg
    └── favicon.svg                  # Icon only, 32×32 viewBox
```

---

## 8. References

- NIP-002: Brand identity implementation plan
- [docs/01_whitepaper.md](../01_whitepaper.md): Tone and positioning reference
